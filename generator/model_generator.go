package generator

import (
	"bytes"
	"fmt"
	apitemplates "go-crud-api-generator/templates"
	"reflect"
	"strings"
	"text/template"
)

func generateModelStructs(refObjectMap map[string]interface{}) []string {
	var structList []string

	for k, v := range refObjectMap {
		structString := renderModelTemplate(v.(map[string]interface{}))
		structString = strings.ReplaceAll(structString, "StructName", k)
		structList = append(structList, structString)
	}
	return structList
}

func renderModelTemplate(structAttributes map[string]interface{}) string {
	var b bytes.Buffer

	// Create a new template
	t := template.Must(template.New("struct").Parse(apitemplates.ModelTemplate))

	// Execute the template
	err := t.Execute(&b, fetchAttributeNameAndType(structAttributes))
	if err != nil {
		panic(err)
	}
	return b.String()
}
func fetchAttributeNameAndType(structAttributes map[string]interface{}) map[string]string {
	sAttrNT := make(map[string]string, 0)
	for k, v := range structAttributes {
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
