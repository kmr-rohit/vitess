/*
Copyright 2022 The Vitess Authors.

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

package main

const _t = `// Code generated by testutil/authztestgen. DO NOT EDIT.

/*
Copyright 2022 The Vitess Authors.

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

package {{ .Package }}

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"vitess.io/vitess/go/vt/vtadmin"
	"vitess.io/vitess/go/vt/vtadmin/cluster"
	"vitess.io/vitess/go/vt/vtadmin/rbac"
	"vitess.io/vitess/go/vt/vtadmin/testutil"
	"vitess.io/vitess/go/vt/vtadmin/vtctldclient/fakevtctldclient"

	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
	vtadminpb "vitess.io/vitess/go/vt/proto/vtadmin"
	vtctldatapb "vitess.io/vitess/go/vt/proto/vtctldata"
)

{{ range .Tests }}
func Test{{ .Method }}(t *testing.T) {
	opts := vtadmin.Options{
		RBAC: &rbac.Config{
			Rules: []*struct{
				Resource string
				Actions  []string
				Subjects []string
				Clusters []string
			}{
				{{- range .Rules }}
				{
					Resource: "{{ .Resource }}",
					Actions:  []string{ {{ range .Actions }}"{{ . }}",{{ end }} },
					Subjects: []string{ {{ range .Subjects }}"{{ . }}",{{ end }} },
					Clusters: []string{ {{ range .Clusters }}"{{ . }}",{{ end }} },
				},
				{{- end }}
			},
		},
	}
	err := opts.RBAC.Reify()
	require.NoError(t, err, "failed to reify authorization rules: %+v", opts.RBAC.Rules)

	api := vtadmin.NewAPI(testClusters(t), opts)
	t.Cleanup(func() {
		if err := api.Close(); err != nil {
			t.Logf("api did not close cleanly: %s", err.Error())
		}
	})

	{{ with $test := . -}}
	{{ range .Cases }}
	t.Run("{{ .Name }}", func(t *testing.T) {
		t.Parallel()
		{{ getActor .Actor }}

		ctx := context.Background()
		if actor != nil {
			ctx = rbac.NewContext(ctx, actor)
		}
		{{ if .IncludeErrorVar }}
		resp, err := api.{{ $test.Method }}(ctx, {{ $test.Request }})
		{{ else }}
		resp, _ := api.{{ $test.Method }}(ctx, {{ $test.Request }})
		{{ end }}
		{{- with $case := . -}}
		{{ range .Assertions }}
		{{- writeAssertion . $test $case }}
		{{ end }}
		{{- end -}}
	})
	{{ end }}
	{{- end -}}
}
{{- end }}

func testClusters(t testing.TB) []*cluster.Cluster {
	configs := []testutil.TestClusterConfig{
		{{ range .Clusters -}}
		{{- with $cluster := . -}}
		{
			Cluster: &vtadminpb.Cluster{
				Id:   "{{ .ID }}",
				Name: "{{ .Name }}",
			},
			VtctldClient: &fakevtctldclient.VtctldClient{
				{{- range .FakeVtctldClientResults }}
				{{ .FieldName }}: {{ .Type }}{
					{{ .Value }}
				},
				{{- end }}
			},
			Tablets: []*vtadminpb.Tablet{
				{{- range .DBTablets }}
				{
					Cluster: &vtadminpb.Cluster{Id: "{{ $cluster.ID }}", Name: "{{ $cluster.Name }}" },
					Tablet: &topodatapb.Tablet{
						Alias: &topodatapb.TabletAlias{
							Cell: "{{ .Tablet.Alias.Cell }}",
							Uid: {{ .Tablet.Alias.Uid }},
						},
					},
				},
				{{- end }}
			},
		},
		{{- end -}}
		{{- end }}
	}

	return testutil.BuildClusters(t, configs...)
}
`
