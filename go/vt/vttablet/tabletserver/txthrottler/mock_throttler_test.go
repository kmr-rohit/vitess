// Code generated by MockGen. DO NOT EDIT.
// Source: vitess.io/vitess/go/vt/throttler (interfaces: Throttler)

// Package txthrottler is a generated GoMock package.
package txthrottler

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"

	discovery "vitess.io/vitess/go/vt/discovery"
	throttlerdata "vitess.io/vitess/go/vt/proto/throttlerdata"
	topodata "vitess.io/vitess/go/vt/proto/topodata"
	throttler "vitess.io/vitess/go/vt/throttler"
)

// MockThrottler is a mock of Throttler interface.
type MockThrottler struct {
	ctrl     *gomock.Controller
	recorder *MockThrottlerMockRecorder
}

// MockThrottlerMockRecorder is the mock recorder for MockThrottler.
type MockThrottlerMockRecorder struct {
	mock *MockThrottler
}

// NewMockThrottler creates a new mock instance.
func NewMockThrottler(ctrl *gomock.Controller) *MockThrottler {
	mock := &MockThrottler{ctrl: ctrl}
	mock.recorder = &MockThrottlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThrottler) EXPECT() *MockThrottlerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockThrottler) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockThrottlerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockThrottler)(nil).Close))
}

// GetConfiguration mocks base method.
func (m *MockThrottler) GetConfiguration() *throttlerdata.Configuration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfiguration")
	ret0, _ := ret[0].(*throttlerdata.Configuration)
	return ret0
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *MockThrottlerMockRecorder) GetConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockThrottler)(nil).GetConfiguration))
}

// Log mocks base method.
func (m *MockThrottler) Log() []throttler.Result {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log")
	ret0, _ := ret[0].([]throttler.Result)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockThrottlerMockRecorder) Log() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockThrottler)(nil).Log))
}

// MaxLag mocks base method.
func (m *MockThrottler) MaxLag(arg0 topodata.TabletType) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxLag", arg0)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// MaxLag indicates an expected call of MaxLag.
func (mr *MockThrottlerMockRecorder) MaxLag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxLag", reflect.TypeOf((*MockThrottler)(nil).MaxLag), arg0)
}

// MaxRate mocks base method.
func (m *MockThrottler) MaxRate() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxRate")
	ret0, _ := ret[0].(int64)
	return ret0
}

// MaxRate indicates an expected call of MaxRate.
func (mr *MockThrottlerMockRecorder) MaxRate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxRate", reflect.TypeOf((*MockThrottler)(nil).MaxRate))
}

// RecordReplicationLag mocks base method.
func (m *MockThrottler) RecordReplicationLag(arg0 time.Time, arg1 *discovery.TabletHealth) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RecordReplicationLag", arg0, arg1)
}

// RecordReplicationLag indicates an expected call of RecordReplicationLag.
func (mr *MockThrottlerMockRecorder) RecordReplicationLag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordReplicationLag", reflect.TypeOf((*MockThrottler)(nil).RecordReplicationLag), arg0, arg1)
}

// ResetConfiguration mocks base method.
func (m *MockThrottler) ResetConfiguration() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetConfiguration")
}

// ResetConfiguration indicates an expected call of ResetConfiguration.
func (mr *MockThrottlerMockRecorder) ResetConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetConfiguration", reflect.TypeOf((*MockThrottler)(nil).ResetConfiguration))
}

// SetMaxRate mocks base method.
func (m *MockThrottler) SetMaxRate(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxRate", arg0)
}

// SetMaxRate indicates an expected call of SetMaxRate.
func (mr *MockThrottlerMockRecorder) SetMaxRate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxRate", reflect.TypeOf((*MockThrottler)(nil).SetMaxRate), arg0)
}

// ThreadFinished mocks base method.
func (m *MockThrottler) ThreadFinished(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ThreadFinished", arg0)
}

// ThreadFinished indicates an expected call of ThreadFinished.
func (mr *MockThrottlerMockRecorder) ThreadFinished(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ThreadFinished", reflect.TypeOf((*MockThrottler)(nil).ThreadFinished), arg0)
}

// Throttle mocks base method.
func (m *MockThrottler) Throttle(arg0 int) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Throttle", arg0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Throttle indicates an expected call of Throttle.
func (mr *MockThrottlerMockRecorder) Throttle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Throttle", reflect.TypeOf((*MockThrottler)(nil).Throttle), arg0)
}

// UpdateConfiguration mocks base method.
func (m *MockThrottler) UpdateConfiguration(arg0 *throttlerdata.Configuration, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfiguration indicates an expected call of UpdateConfiguration.
func (mr *MockThrottlerMockRecorder) UpdateConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfiguration", reflect.TypeOf((*MockThrottler)(nil).UpdateConfiguration), arg0, arg1)
}
