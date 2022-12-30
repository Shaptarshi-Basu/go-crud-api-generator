package specparser


func parseRefObjectMap(refObjectMap map[string]interface{}) {
	for k, v := range refObjectMap {
		t = reflect.Append(t, reflect.StructField{
			Name: k,
			Type: reflect.TypeOf(v),
		})
	}

}