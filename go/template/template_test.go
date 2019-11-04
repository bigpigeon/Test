/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package template

import (
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
