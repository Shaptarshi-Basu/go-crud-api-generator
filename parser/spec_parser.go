package specparser

import (
	"bytes"
	"strings"
	"text/template"
)

func ParseRefObjectMap(refObjectMap map[string]interface{}) []string {
	var structList []string

	for k, v := range refObjectMap {
		structString := createTemplateStruct(v.(map[string]interface{}))
		structString = strings.ReplaceAll(structString, "StructName", k)
		structList = append(structList, structString)
	}
	return structList
}
func createTemplateStruct(structAttributes map[string]interface{}) string {
	const tmpl = `
	type StructName struct {
		{{- range $key, $value := . }}
		{{ $key }} {{ $value }}
		{{- end }}
	}
	`

	var b bytes.Buffer

	// Create a new template
	t := template.Must(template.New("struct").Parse(tmpl))

	// Execute the template
	err := t.Execute(&b, structAttributes)
	if err != nil {
		panic(err)
	}
	return b.String()
}
