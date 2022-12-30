package specparser

import (
	"fmt"
	"reflect"
	"strings"
)

func ParseRefObjectMap(refObjectMap map[string]interface{}) {
	structMap := make(map[string]string)

	var fieldNames []string
	var fieldTypes []reflect.Type
	for k, v := range refObjectMap {
		fieldNames, fieldTypes = fetchAllFieldNamesAndTypes(v.(map[string]interface{}))

		// Create a new struct type
		fields := make([]reflect.StructField, len(fieldNames))
		for i := 0; i < len(fieldNames); i++ {
			fields[i] = reflect.StructField{
				Name: fieldNames[i],
				Type: fieldTypes[i],
			}
		}
		t := reflect.StructOf(fields)
		structMap[k] = fmt.Sprintf("type %s %s", k, t)
		fmt.Println(structMap[k])

	}
}
func fetchAllFieldNamesAndTypes(structAttributes map[string]interface{}) (fieldNames []string, fieldTypes []reflect.Type) {
	for k, v := range structAttributes {
		fieldNames = append(fieldNames, strings.Title(k))
		fieldTypes = append(fieldTypes, reflect.TypeOf(v))
	}
	return
}
