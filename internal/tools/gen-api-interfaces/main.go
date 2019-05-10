package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud-v2/internal/define"
	"github.com/sacloud/libsacloud-v2/internal/tools"
)

const destination = "sacloud/zz_apis.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-interfaces: ")
}

func main() {
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), destination),
		Template:   tmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", filepath.Join(destination))
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-interfaces'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "context" }}
	{{ . }}
{{- end }}
)

{{ range . }}

/************************************************* 
* {{.TypeName}}API
*************************************************/

// {{ .TypeName }}API is interface for operate {{ .TypeName }} resource
type {{ .TypeName }}API interface {
{{ range .AllOperations }}
	{{ .MethodName }}(ctx context.Context{{ range .AllArguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} 
{{- end -}}
}
{{ end }}
`
