/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package template

import (
	"errors"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"text/template"
)

func TestTextTemplate(t *testing.T) {
	temp := `
dear {{.name}}:
{{/*
some comment
*/}}
`
	tree, err := template.New("xxx").Parse(temp)
	require.NoError(t, err)
	err = tree.Execute(os.Stdout, map[string]interface{}{
		"name": "jia",
	})
	require.NoError(t, err)
}

func TestTextElseTemplate(t *testing.T) {
	temp := `
Val:
{{ if .name -}}
dear {{.name}}:
{{- else -}}
dear default
{{- end}}
`
	tree, err := template.New("xxx").Parse(temp)
	require.NoError(t, err)
	err = tree.Execute(os.Stdout, map[string]interface{}{"name": "jia"})
	require.NoError(t, err)

	err = tree.Execute(os.Stdout, map[string]interface{}{})
	require.NoError(t, err)
}

func TestOrCondition(t *testing.T) {
	temp := `
{{- if (or .notExist .notExist2 .name) }}
dear: {{.name}}
{{- end }}
`
	tree, err := template.New("xxx").Parse(temp)
	require.NoError(t, err)
	err = tree.Execute(os.Stdout, map[string]interface{}{"name": "jia"})
	require.NoError(t, err)

	err = tree.Execute(os.Stdout, map[string]interface{}{})
	require.NoError(t, err)
}

func TestServeralPipelineParams(t *testing.T) {
	const temp = `{{- define "userlist" }}
total: {{.Total}}
current: {{.Current}}
{{- end}}
start
  {{- template "userlist" dict "Total" .total "Current" .current}}
`

	tree := template.New("xxx").Funcs(template.FuncMap{
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
	})
	_, err := tree.Parse(temp)
	require.NoError(t, err)
	err = tree.Execute(os.Stdout, map[string]interface{}{"total": 123, "current": 2})
	require.NoError(t, err)
}
