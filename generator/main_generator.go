package generator

import (
	"bytes"
	apitemplates "go-crud-api-generator/templates"
	"text/template"
)

func (api *apiGenrerator) generateMainFunc() string {

	var b bytes.Buffer
	// Create a new template
	t := template.Must(template.New("main").Parse(apitemplates.MainTemplate))
	err := t.Execute(&b, api.spec.Paths)
	if err != nil {
		panic(err)
	}
	return b.String()
}
