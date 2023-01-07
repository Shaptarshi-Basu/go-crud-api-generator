package generator

import (
	"bytes"
	apitemplates "go-crud-api-generator/templates"
	"text/template"
)

func (api *apiGenrerator) generateHandlers() string {
	var b bytes.Buffer

	t := template.Must(template.New("main").Parse(apitemplates.HandlerFuncTemplate))

	// Execute the template
	err := t.Execute(&b, api.spec.Paths)
	if err != nil {
		panic(err)
	}
	return b.String()
}
