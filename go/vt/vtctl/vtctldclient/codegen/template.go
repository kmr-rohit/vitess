/*
Copyright 2021 The Vitess Authors.

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

import "text/template"

const tmplStr = `// Code generated by {{ .ClientName }}-generator. DO NOT EDIT.

/*
Copyright 2021 The Vitess Authors.

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

package {{ .PackageName }}

import (
	"context"

	"google.golang.org/grpc"
	{{ if not .Local -}}
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	{{- end }}

	{{ range .Imports -}}
	{{ if ne .Alias "" }}{{ .Alias }} {{ end }}"{{ .Path }}"
	{{- end }}
)
{{ range .Methods }}
// {{ .Name }} is part of the vtctlservicepb.VtctldClient interface.
func (client *{{ $.Type }}) {{ .Name }}(ctx context.Context, {{ .Param.Name }} {{ .Param.Type }}, opts ...grpc.CallOption) ({{ .Result.Type }}, error) {
	{{- if not $.Local -}}
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	{{ end -}}
	return client.{{ if $.Local }}s{{ else }}c{{ end }}.{{ .Name }}(ctx, in{{ if not $.Local }}, opts...{{ end }})
}
{{ end }}`

var tmpl = template.Must(template.New("vtctldclient-generator").Parse(tmplStr))
