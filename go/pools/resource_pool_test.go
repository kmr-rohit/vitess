/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pools

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"context"

	"vitess.io/vitess/go/sync2"
)

var lastID, count, closeCount sync2.AtomicInt64
var waitStarts []time.Time

type TestResource struct {
	num       int64
	closed    bool
	setting   []string
	failApply bool
}

func (tr *TestResource) ApplySettings(_ context.Context, settings []string) error {
	if tr.failApply {
		return fmt.Errorf("ApplySettings failed")
	}
	tr.setting = settings
	return nil
}

func (tr *TestResource) IsSettingsApplied() bool {
	return len(tr.setting) > 0
}

func (tr *TestResource) IsSameSetting(settings []string) bool {
	return compareStringSlice(tr.setting, settings)
}

func compareStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aVal := range a {
		if aVal != b[i] {
			return false
		}
	}
	return true
}

func (tr *TestResource) Close() {
	if !tr.closed {
		count.Add(-1)
		closeCount.Add(1)
		tr.closed = true
	}
}

func logWait(start time.Time) {
	waitStarts = append(waitStarts, start)
}

func PoolFactory(context.Context) (Resource, error) {
	count.Add(1)
	return &TestResource{num: lastID.Add(1)}, nil
}

func FailFactory(context.Context) (Resource, error) {
	return nil, errors.New("Failed")
}

func SlowFailFactory(context.Context) (Resource, error) {
	time.Sleep(10 * time.Millisecond)
	return nil, errors.New("Failed")
}

func DisallowSettingsFactory(context.Context) (Resource, error) {
	count.Add(1)
	return &TestResource{num: lastID.Add(1), failApply: true}, nil
}

func TestOpen(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	waitStarts = waitStarts[:0]

	p := NewResourcePool(PoolFactory, 6, 6, time.Second, logWait, nil, 0)
	p.SetCapacity(5)
	var resources [10]Resource
	var r Resource
	var err error

	// Test Get
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			r, err = p.Get(ctx, nil)
		} else {
			r, err = p.Get(ctx, []string{"foo"})
		}
		require.NoError(t, err)
		resources[i] = r
		assert.EqualValues(t, 5-i-1, p.Available())
		assert.Zero(t, p.WaitCount())
		assert.Zero(t, len(waitStarts))
		assert.Zero(t, p.WaitTime())
		assert.EqualValues(t, i+1, lastID.Get())
		assert.EqualValues(t, i+1, count.Get())
	}

	// Test that Get waits
	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			if i%2 == 0 {
				r, err = p.Get(ctx, nil)
			} else {
				r, err = p.Get(ctx, []string{"foo"})
			}
			require.NoError(t, err)
			resources[i] = r
		}
		for i := 0; i < 5; i++ {
			p.Put(resources[i])
		}
		ch <- true
	}()
	for i := 0; i < 5; i++ {
		// Sleep to ensure the goroutine waits
		time.Sleep(10 * time.Millisecond)
		p.Put(resources[i])
	}
	<-ch
	assert.EqualValues(t, 5, p.WaitCount())
	assert.Equal(t, 5, len(waitStarts))
	// verify start times are monotonic increasing
	for i := 1; i < len(waitStarts); i++ {
		if waitStarts[i].Before(waitStarts[i-1]) {
			t.Errorf("Expecting monotonic increasing start times")
		}
	}
	assert.NotZero(t, p.WaitTime())
	assert.EqualValues(t, 5, lastID.Get())
	// Test Close resource
	r, err = p.Get(ctx, nil)
	require.NoError(t, err)
	r.Close()
	// A nil Put should cause the resource to be reopened.
	p.Put(nil)
	assert.EqualValues(t, 5, count.Get())
	assert.EqualValues(t, 6, lastID.Get())

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			r, err = p.Get(ctx, nil)
		} else {
			r, err = p.Get(ctx, []string{"foo"})
		}
		require.NoError(t, err)
		resources[i] = r
	}
	for i := 0; i < 5; i++ {
		p.Put(resources[i])
	}
	assert.EqualValues(t, 5, count.Get())
	assert.EqualValues(t, 6, lastID.Get())

	// SetCapacity
	p.SetCapacity(3)
	assert.EqualValues(t, 3, count.Get())
	assert.EqualValues(t, 6, lastID.Get())
	assert.EqualValues(t, 3, p.Capacity())
	assert.EqualValues(t, 3, p.Available())

	p.SetCapacity(6)
	assert.EqualValues(t, 6, p.Capacity())
	assert.EqualValues(t, 6, p.Available())

	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			r, err = p.Get(ctx, nil)
		} else {
			r, err = p.Get(ctx, []string{"foo"})
		}
		require.NoError(t, err)
		resources[i] = r
	}
	for i := 0; i < 6; i++ {
		p.Put(resources[i])
	}
	assert.EqualValues(t, 6, count.Get())
	assert.EqualValues(t, 9, lastID.Get())

	// Close
	p.Close()
	assert.EqualValues(t, 0, p.Capacity())
	assert.EqualValues(t, 0, p.Available())
	assert.EqualValues(t, 0, count.Get())
}

func TestShrinking(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	waitStarts = waitStarts[:0]

	p := NewResourcePool(PoolFactory, 5, 5, time.Second, logWait, nil, 0)
	var resources [10]Resource
	// Leave one empty slot in the pool
	for i := 0; i < 4; i++ {
		var r Resource
		var err error
		if i%2 == 0 {
			r, err = p.Get(ctx, nil)
		} else {
			r, err = p.Get(ctx, []string{"foo"})
		}
		require.NoError(t, err)
		resources[i] = r
	}
	done := make(chan bool)
	go func() {
		p.SetCapacity(3)
		done <- true
	}()
	expected := `{"Capacity": 3, "Available": 0, "Active": 4, "InUse": 4, "MaxCapacity": 5, "WaitCount": 0, "WaitTime": 0, "IdleTimeout": 1000000000, "IdleClosed": 0, "Exhausted": 0}`
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		stats := p.StatsJSON()
		if stats != expected {
			if i == 9 {
				t.Errorf(`expecting '%s', received '%s'`, expected, stats)
			}
		}
	}
	// There are already 2 resources available in the pool.
	// So, returning one should be enough for SetCapacity to complete.
	p.Put(resources[3])
	<-done
	// Return the rest of the resources
	for i := 0; i < 3; i++ {
		p.Put(resources[i])
	}
	stats := p.StatsJSON()
	expected = `{"Capacity": 3, "Available": 3, "Active": 3, "InUse": 0, "MaxCapacity": 5, "WaitCount": 0, "WaitTime": 0, "IdleTimeout": 1000000000, "IdleClosed": 0, "Exhausted": 0}`
	assert.Equal(t, expected, stats)
	assert.EqualValues(t, 3, count.Get())

	// Ensure no deadlock if SetCapacity is called after we start
	// waiting for a resource
	var err error
	for i := 0; i < 3; i++ {
		var r Resource
		if i%2 == 0 {
			r, err = p.Get(ctx, nil)
		} else {
			r, err = p.Get(ctx, []string{"foo"})
		}
		require.NoError(t, err)
		resources[i] = r
	}
	// This will wait because pool is empty
	go func() {
		r, err := p.Get(ctx, nil)
		require.NoError(t, err)
		p.Put(r)
		done <- true
	}()

	// This will also wait
	go func() {
		p.SetCapacity(2)
		done <- true
	}()
	time.Sleep(10 * time.Millisecond)

	// This should not hang
	for i := 0; i < 3; i++ {
		p.Put(resources[i])
	}
	<-done
	<-done
	assert.EqualValues(t, 2, p.Capacity())
	assert.EqualValues(t, 2, p.Available())
	assert.EqualValues(t, 1, p.WaitCount())
	assert.EqualValues(t, p.WaitCount(), len(waitStarts))
	assert.EqualValues(t, 2, count.Get())

	// Test race condition of SetCapacity with itself
	p.SetCapacity(3)
	for i := 0; i < 3; i++ {
		var r Resource
		var err error
		if i%2 == 0 {
			r, err = p.Get(ctx, nil)
		} else {
			r, err = p.Get(ctx, []string{"foo"})
		}
		require.NoError(t, err)
		resources[i] = r
	}
	// This will wait because pool is empty
	go func() {
		r, err := p.Get(ctx, nil)
		require.NoError(t, err)
		p.Put(r)
		done <- true
	}()
	time.Sleep(10 * time.Millisecond)

	// This will wait till we Put
	go p.SetCapacity(2)
	time.Sleep(10 * time.Millisecond)
	go p.SetCapacity(4)
	time.Sleep(10 * time.Millisecond)

	// This should not hang
	for i := 0; i < 3; i++ {
		p.Put(resources[i])
	}
	<-done

	err = p.SetCapacity(-1)
	if err == nil {
		t.Errorf("Expecting error")
	}
	err = p.SetCapacity(255555)
	if err == nil {
		t.Errorf("Expecting error")
	}

	assert.EqualValues(t, 4, p.Capacity())
	assert.EqualValues(t, 4, p.Available())
}

func TestClosing(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	p := NewResourcePool(PoolFactory, 5, 5, time.Second, logWait, nil, 0)
	var resources [10]Resource
	for i := 0; i < 5; i++ {
		var r Resource
		var err error
		if i%2 == 0 {
			r, err = p.Get(ctx, nil)
		} else {
			r, err = p.Get(ctx, []string{"foo"})
		}
		require.NoError(t, err)
		resources[i] = r
	}
	ch := make(chan bool)
	go func() {
		p.Close()
		ch <- true
	}()

	// Wait for goroutine to call Close
	time.Sleep(10 * time.Millisecond)
	stats := p.StatsJSON()
	expected := `{"Capacity": 0, "Available": 0, "Active": 5, "InUse": 5, "MaxCapacity": 5, "WaitCount": 0, "WaitTime": 0, "IdleTimeout": 1000000000, "IdleClosed": 0, "Exhausted": 1}`
	assert.Equal(t, expected, stats)

	// Put is allowed when closing
	for i := 0; i < 5; i++ {
		p.Put(resources[i])
	}

	// Wait for Close to return
	<-ch

	stats = p.StatsJSON()
	expected = `{"Capacity": 0, "Available": 0, "Active": 0, "InUse": 0, "MaxCapacity": 5, "WaitCount": 0, "WaitTime": 0, "IdleTimeout": 1000000000, "IdleClosed": 0, "Exhausted": 1}`
	assert.Equal(t, expected, stats)
	assert.EqualValues(t, 5, lastID.Get())
	assert.EqualValues(t, 0, count.Get())
}

func TestReopen(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	refreshCheck := func() (bool, error) {
		return true, nil
	}
	p := NewResourcePool(PoolFactory, 5, 5, time.Second, logWait, refreshCheck, 500*time.Millisecond)
	var resources [10]Resource
	for i := 0; i < 5; i++ {
		var r Resource
		var err error
		if i%2 == 0 {
			r, err = p.Get(ctx, nil)
		} else {
			r, err = p.Get(ctx, []string{"foo"})
		}
		require.NoError(t, err)
		resources[i] = r
	}

	time.Sleep(10 * time.Millisecond)
	stats := p.StatsJSON()
	expected := `{"Capacity": 5, "Available": 0, "Active": 5, "InUse": 5, "MaxCapacity": 5, "WaitCount": 0, "WaitTime": 0, "IdleTimeout": 1000000000, "IdleClosed": 0, "Exhausted": 1}`
	assert.Equal(t, expected, stats)

	time.Sleep(650 * time.Millisecond)
	for i := 0; i < 5; i++ {
		p.Put(resources[i])
	}
	time.Sleep(50 * time.Millisecond)
	stats = p.StatsJSON()
	expected = `{"Capacity": 5, "Available": 5, "Active": 0, "InUse": 0, "MaxCapacity": 5, "WaitCount": 0, "WaitTime": 0, "IdleTimeout": 1000000000, "IdleClosed": 0, "Exhausted": 1}`
	assert.Equal(t, expected, stats)
	assert.EqualValues(t, 5, lastID.Get())
	assert.EqualValues(t, 0, count.Get())
}

func TestIdleTimeout(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	p := NewResourcePool(PoolFactory, 1, 1, 10*time.Millisecond, logWait, nil, 0)
	defer p.Close()

	r, err := p.Get(ctx, nil)
	require.NoError(t, err)
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 0, p.IdleClosed())

	p.Put(r)
	assert.EqualValues(t, 1, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 0, p.IdleClosed())

	time.Sleep(15 * time.Millisecond)
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	r, err = p.Get(ctx, nil)
	require.NoError(t, err)
	assert.EqualValues(t, 2, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	// sleep to let the idle closer run while all resources are in use
	// then make sure things are still as we expect
	time.Sleep(15 * time.Millisecond)
	assert.EqualValues(t, 2, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	p.Put(r)
	r, err = p.Get(ctx, nil)
	require.NoError(t, err)
	assert.EqualValues(t, 2, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	// the idle close thread wakes up every 1/100 of the idle time, so ensure
	// the timeout change applies to newly added resources
	p.SetIdleTimeout(1000 * time.Millisecond)
	p.Put(r)

	time.Sleep(15 * time.Millisecond)
	assert.EqualValues(t, 2, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	// Get and Put to refresh timeUsed
	r, err = p.Get(ctx, nil)
	require.NoError(t, err)
	p.Put(r)
	p.SetIdleTimeout(10 * time.Millisecond)
	time.Sleep(15 * time.Millisecond)
	assert.EqualValues(t, 3, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 2, p.IdleClosed())
}

func TestIdleTimeoutWithSettings(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	p := NewResourcePool(PoolFactory, 1, 1, 10*time.Millisecond, logWait, nil, 0)
	defer p.Close()

	r, err := p.Get(ctx, []string{"foo", "bar"})
	require.NoError(t, err)
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 0, p.IdleClosed())

	p.Put(r)
	assert.EqualValues(t, 1, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 0, p.IdleClosed())

	time.Sleep(15 * time.Millisecond)
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	r, err = p.Get(ctx, []string{"foo", "bar"})
	require.NoError(t, err)
	assert.EqualValues(t, 2, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	// sleep to let the idle closer run while all resources are in use
	// then make sure things are still as we expect
	time.Sleep(15 * time.Millisecond)
	assert.EqualValues(t, 2, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	p.Put(r)
	r, err = p.Get(ctx, []string{"foo", "bar"})
	require.NoError(t, err)
	assert.EqualValues(t, 2, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	// the idle close thread wakes up every 1/100 of the idle time, so ensure
	// the timeout change applies to newly added resources
	p.SetIdleTimeout(1000 * time.Millisecond)
	p.Put(r)

	time.Sleep(15 * time.Millisecond)
	assert.EqualValues(t, 2, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 1, p.IdleClosed())

	// Get and Put to refresh timeUsed
	r, err = p.Get(ctx, []string{"foo", "bar"})
	require.NoError(t, err)
	p.Put(r)
	p.SetIdleTimeout(10 * time.Millisecond)
	time.Sleep(15 * time.Millisecond)
	assert.EqualValues(t, 3, lastID.Get())
	assert.EqualValues(t, 1, count.Get())
	assert.EqualValues(t, 2, p.IdleClosed())
}

func TestIdleTimeoutCreateFail(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	p := NewResourcePool(PoolFactory, 1, 1, 10*time.Millisecond, logWait, nil, 0)
	defer p.Close()
	for _, setting := range [][]string{nil, {"foo"}} {
		r, err := p.Get(ctx, setting)
		require.NoError(t, err)
		// Change the factory before putting back
		// to prevent race with the idle closer, who will
		// try to use it.
		p.factory = FailFactory
		p.Put(r)
		time.Sleep(15 * time.Millisecond)
		assert.Zero(t, p.Active())

		// reset factory for next run.
		p.factory = PoolFactory
	}
}

func TestCreateFail(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	p := NewResourcePool(FailFactory, 5, 5, time.Second, logWait, nil, 0)
	defer p.Close()

	for _, setting := range [][]string{nil, {"foo"}} {
		if _, err := p.Get(ctx, setting); err.Error() != "Failed" {
			t.Errorf("Expecting Failed, received %v", err)
		}
		stats := p.StatsJSON()
		expected := `{"Capacity": 5, "Available": 5, "Active": 0, "InUse": 0, "MaxCapacity": 5, "WaitCount": 0, "WaitTime": 0, "IdleTimeout": 1000000000, "IdleClosed": 0, "Exhausted": 0}`
		assert.Equal(t, expected, stats)
	}
}

func TestCreateFailOnPut(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	p := NewResourcePool(PoolFactory, 5, 5, time.Second, logWait, nil, 0)
	defer p.Close()

	for _, setting := range [][]string{nil, {"foo"}} {
		_, err := p.Get(ctx, setting)
		require.NoError(t, err)

		// change factory to fail the put.
		p.factory = FailFactory
		p.Put(nil)
		assert.Zero(t, p.Active())

		// change back for next iteration.
		p.factory = PoolFactory
	}
}

func TestSlowCreateFail(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	p := NewResourcePool(SlowFailFactory, 2, 2, time.Second, logWait, nil, 0)
	defer p.Close()
	ch := make(chan bool)
	for _, setting := range [][]string{nil, {"foo"}} {
		// The third Get should not wait indefinitely
		for i := 0; i < 3; i++ {
			go func() {
				p.Get(ctx, setting)
				ch <- true
			}()
		}
		for i := 0; i < 3; i++ {
			<-ch
		}
		assert.EqualValues(t, 2, p.Available())
	}
}

func TestTimeout(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	p := NewResourcePool(PoolFactory, 1, 1, time.Second, logWait, nil, 0)
	defer p.Close()

	// take the only connection available
	r, err := p.Get(ctx, nil)
	require.NoError(t, err)

	for _, setting := range [][]string{nil, {"foo"}} {
		// trying to get the connection without a timeout.
		newctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
		_, err = p.Get(newctx, setting)
		cancel()
		assert.EqualError(t, err, "resource pool timed out")

	}

	// put the connection take was taken initially.
	p.Put(r)
}

func TestExpired(t *testing.T) {
	lastID.Set(0)
	count.Set(0)
	p := NewResourcePool(PoolFactory, 1, 1, time.Second, logWait, nil, 0)
	defer p.Close()

	for _, setting := range [][]string{nil, {"foo"}} {
		// expired context
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-1*time.Second))
		_, err := p.Get(ctx, setting)
		cancel()
		require.EqualError(t, err, "resource pool context already expired")
	}
}

func TestMultiSettings(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	waitStarts = waitStarts[:0]

	p := NewResourcePool(PoolFactory, 5, 5, time.Second, logWait, nil, 0)
	var resources [10]Resource
	var r Resource
	var err error

	settings := [][]string{nil, {"foo"}, {"bar"}, {"bar"}, {"foo"}}

	// Test Get
	for i := 0; i < 5; i++ {
		r, err = p.Get(ctx, settings[i])
		require.NoError(t, err)
		resources[i] = r
		assert.EqualValues(t, 5-i-1, p.Available())
		assert.Zero(t, p.WaitCount())
		assert.Zero(t, len(waitStarts))
		assert.Zero(t, p.WaitTime())
		assert.EqualValues(t, i+1, lastID.Get())
		assert.EqualValues(t, i+1, count.Get())
	}

	// Test that Get waits
	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			r, err = p.Get(ctx, settings[i])
			require.NoError(t, err)
			resources[i] = r
		}
		for i := 0; i < 5; i++ {
			p.Put(resources[i])
		}
		ch <- true
	}()
	for i := 0; i < 5; i++ {
		// Sleep to ensure the goroutine waits
		time.Sleep(10 * time.Millisecond)
		p.Put(resources[i])
	}
	<-ch
	assert.EqualValues(t, 5, p.WaitCount())
	assert.Equal(t, 5, len(waitStarts))
	// verify start times are monotonic increasing
	for i := 1; i < len(waitStarts); i++ {
		if waitStarts[i].Before(waitStarts[i-1]) {
			t.Errorf("Expecting monotonic increasing start times")
		}
	}
	assert.NotZero(t, p.WaitTime())
	assert.EqualValues(t, 5, lastID.Get())

	// Close
	p.Close()
	assert.EqualValues(t, 0, p.Capacity())
	assert.EqualValues(t, 0, p.Available())
	assert.EqualValues(t, 0, count.Get())
}

func TestMultiSettingsWithClosures(t *testing.T) {
	ctx := context.Background()
	lastID.Set(0)
	count.Set(0)
	closeCount.Set(0)

	p := NewResourcePool(PoolFactory, 5, 5, time.Second, logWait, nil, 0)
	var resources [10]Resource
	var r Resource
	var err error

	settings := [][]string{nil, {"foo"}, {"bar"}, {"bar"}, {"foo"}}

	// Test Get
	for i := 0; i < 5; i++ {
		r, err = p.Get(ctx, settings[i])
		require.NoError(t, err)
		resources[i] = r
		assert.EqualValues(t, 5-i-1, p.Available())
		assert.EqualValues(t, i+1, lastID.Get())
		assert.EqualValues(t, i+1, count.Get())
	}

	// Put all of them back
	for i := 0; i < 5; i++ {
		p.Put(resources[i])
	}

	// Getting all with same settings.
	for i := 0; i < 5; i++ {
		r, err = p.Get(ctx, settings[1]) // {foo}
		require.NoError(t, err)
		p.Put(r)
	}
	assert.EqualValues(t, 2, closeCount.Get()) // when setting was {bar} and getting for {foo}
	assert.EqualValues(t, 5, p.Available())
	assert.EqualValues(t, 7, lastID.Get())
	assert.EqualValues(t, 5, count.Get())

	// Close
	p.Close()
	assert.EqualValues(t, 0, p.Capacity())
	assert.EqualValues(t, 0, p.Available())
	assert.EqualValues(t, 0, count.Get())
}

func TestApplySettingsFailure(t *testing.T) {
	ctx := context.Background()
	var resources []Resource
	var r Resource
	var err error

	p := NewResourcePool(PoolFactory, 5, 5, time.Second, logWait, nil, 0)
	defer p.Close()

	settings := [][]string{nil, {"foo"}, {"bar"}, {"bar"}, {"foo"}}
	// get the resource and mark for failure
	for i := 0; i < 5; i++ {
		r, err = p.Get(ctx, settings[i])
		require.NoError(t, err)
		r.(*TestResource).failApply = true
		resources = append(resources, r)
	}
	// put them back
	for _, r = range resources {
		p.Put(r)
	}

	// any new connection created will fail to apply settings
	p.factory = DisallowSettingsFactory

	// Get the resource with "foo" settings
	// For an applied connection if the settings are same it will be returned as-is.
	// Otherwise, will fail to get the resource.
	var failCount int
	resources = nil
	for i := 0; i < 5; i++ {
		r, err = p.Get(ctx, settings[1])
		if err != nil {
			failCount++
			assert.EqualError(t, err, "ApplySettings failed")
			continue
		}
		resources = append(resources, r)
	}
	// put them back
	for _, r = range resources {
		p.Put(r)
	}
	require.Equal(t, 3, failCount)

	// should be able to get all the resource with no settings
	resources = nil
	for i := 0; i < 5; i++ {
		r, err = p.Get(ctx, nil)
		require.NoError(t, err)
		resources = append(resources, r)
	}
	// put them back
	for _, r = range resources {
		p.Put(r)
	}
}
