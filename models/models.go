package models

import (
	"encoding/json"
	"go-crud-api-generator/util"
	"io/ioutil"
	"strings"
)

type Paths map[string][]Path

type Spec struct {
	Paths Paths                  `json:"paths"`
	Refs  map[string]interface{} `json:"refs"`
}

type Path struct {
	Method    string `json:"method"`
	FuncName  string
	Request   interface{}   `json:"request"`
	Responses []interface{} `json:"responses"`
}

func (spec *Spec) refineRouteDetails() {
	for k, paths := range spec.Paths {
		var editedPaths []Path
		for _, path := range paths {
			path.FuncName = strings.Title(util.CovertPathtoCamelCaseMethodName(k, strings.ToLower(path.Method)))
			path.Method = strings.ToUpper(path.Method)
			editedPaths = append(editedPaths, path)
		}
		spec.Paths[k] = editedPaths
	}
}

func (spec *Spec) FetchDetailsFromSpecFile(filePath string) error {
	byt, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(byt, &spec); err != nil {
		return err
	}
	spec.refineRouteDetails()
	return nil
}
