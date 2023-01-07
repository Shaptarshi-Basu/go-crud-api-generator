package generator

import (
	"bytes"
	"fmt"
	apitemplates "go-crud-api-generator/templates"
	"text/template"
)

type modFile struct {
	ProjectName string
	GoVersion   string
}

func (api *apiGenrerator) generateModFile() string {

	var b bytes.Buffer
	// Create a new template
	t := template.Must(template.New("main").Parse(apitemplates.ModeFileTemplate))
	modFile := modFile{ProjectName: api.env.ProjectName, GoVersion: api.env.GoVersion}
	fmt.Print("The mod file is ", modFile)
	err := t.Execute(&b, modFile)
	if err != nil {
		panic(err)
	}
	return b.String()
}
