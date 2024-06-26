// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package templates

import (
	"io"
	"text/template"

	"github.com/uyuni-project/uyuni-tools/shared/types"
)

const inspectTemplate = `#!/bin/bash
# inspect.sh, generated by mgradm
{{- range .Param }}
	{{- if or (not $.ProxyHost) (and $.ProxyHost .Proxy ) }}
echo "{{ .Variable }}=$({{ .CLI }})" >> {{ $.OutputFile }}
{{- end }}
{{- end }}
exit 0
`

// InspectTemplateData represents information used to create inspect script.
type InspectTemplateData struct {
	Param      []types.InspectData
	OutputFile string
	ProxyHost  bool
}

// Render will create inspect script.
func (data InspectTemplateData) Render(wr io.Writer) error {
	t := template.Must(template.New("inspect").Parse(inspectTemplate))
	return t.Execute(wr, data)
}
