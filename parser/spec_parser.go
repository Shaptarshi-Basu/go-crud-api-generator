package specparser

import (
	"bytes"
	"fmt"
	"go-crud-api-generator/models"
	apitemplates "go-crud-api-generator/templates"
	"go-crud-api-generator/util"
	"reflect"
	"strings"
	"text/template"
)

type ApiInfo struct {
	PathsMap map[string][]models.Path
}

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
	err := t.Execute(&b, fetchAttributeNameAndType(structAttributes))
	if err != nil {
		panic(err)
	}
	return b.String()
}

func (api *ApiInfo) CreateMainFunc(pathsMap map[string][]models.Path) string {

	var b bytes.Buffer
	// Create a new template
	t := template.Must(template.New("main").Parse(apitemplates.MainTemplate))

	// Execute the template
	pathsMap = fetchRouteDetailsForMain(pathsMap)
	api.PathsMap = pathsMap
	err := t.Execute(&b, pathsMap)
	if err != nil {
		panic(err)
	}
	return b.String()
}

func (api *ApiInfo) CreateHandlerFunc() string {
	var b bytes.Buffer

	t := template.Must(template.New("main").Parse(apitemplates.HandlerFuncTemplate))

	// Execute the template
	err := t.Execute(&b, api.PathsMap)
	if err != nil {
		panic(err)
	}
	return b.String()
}

func fetchRouteDetailsForMain(pathsMap map[string][]models.Path) map[string][]models.Path {
	fmt.Printf("Map supplied was %+v", pathsMap)
	for k, paths := range pathsMap {
		var editedPaths []models.Path
		for _, path := range paths {
			path.FuncName = strings.Title(util.CovertPathtoCamelCaseMethodName(k, strings.ToLower(path.Method)))
			path.Method = strings.ToUpper(path.Method)
			editedPaths = append(editedPaths, path)
		}
		pathsMap[k] = editedPaths
	}
	fmt.Println()
	fmt.Printf("Processed map is %+v", pathsMap)
	fmt.Printf("Processed map len is %+v", len(pathsMap))
	return pathsMap
}

func fetchAttributeNameAndType(structAttributes map[string]interface{}) map[string]string {
	sAttrNT := make(map[string]string, 0)
	for k, v := range structAttributes {
		fmt.Print("value is ", v)
		switch t := reflect.TypeOf(v); t.Kind() {
		case reflect.String:
			if strings.HasPrefix(v.(string), "ref.") {
				sAttrNT[k] = strings.TrimPrefix(v.(string), "ref.")
			} else {
				sAttrNT[k] = "string"
			}
		case reflect.Bool:
			sAttrNT[k] = "bool"
		case reflect.Slice:
			if sVal, ok := v.([]interface{})[0].(string); ok {
				if strings.HasPrefix(sVal, "ref.") {
					sAttrNT[k] = fmt.Sprintf("[]%s", strings.TrimPrefix(sVal, "ref."))
				} else {
					sAttrNT[k] = fmt.Sprintf("[]%s", "string")
				}
			} else {
				sAttrNT[k] = fmt.Sprintf("[]%T", v.([]interface{})[0])
			}

		case reflect.Map:
			sAttrNT[k] = fmt.Sprintf("%T", v)
		case reflect.Int64:
			sAttrNT[k] = "int"
		case reflect.Float64:
			sAttrNT[k] = "float64"
		default:
			sAttrNT[k] = t.String()
		}
	}
	return sAttrNT
}
