package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud/v2/internal/define"
	"github.com/sacloud/libsacloud/v2/internal/tools"
)

const destination = "sacloud/zz_result.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-result: ")
}

func main() {
	outputPath := destination
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), outputPath),
		Template:   tmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", outputPath)

}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-result'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "github.com/sacloud/libsacloud/v2/sacloud/types" }}
	{{ . }}
{{- end }}
)

{{- range . }}
{{- range .Operations -}}

{{ if .HasResponseEnvelope }}
// {{ .ResultTypeName }} represents the Result of API 
type {{ .ResultTypeName }} struct {
{{- if .IsResponsePlural -}}
	Total       int        ` + "`" + `json:",omitempty"` + "`" + ` // Total count of target resources
	From        int        ` + "`" + `json:",omitempty"` + "`" + ` // Current page number
	Count       int        ` + "`" + `json:",omitempty"` + "`" + ` // Count of current page
{{ else }}
	IsOk    bool  ` + "`" + `json:"is_ok,omitempty"` + "`" + ` // is_ok
{{ end }}
{{ if .IsResponseSingular }}
	{{- range .ResponsePayloads}}
	{{.PayloadName}} {{.TypeName}}
	{{- end }}
{{- else if .IsResponsePlural -}}
	{{- range .ResponsePayloads}}
	{{.PayloadName}} []{{.TypeName}}
	{{- end }}
{{ end }}
}
{{ end }}

{{- end -}}
{{- end -}}
`