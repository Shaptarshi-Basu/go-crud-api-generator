package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-crud-api-generator/models"
	specparser "go-crud-api-generator/parser"
	"io/ioutil"
)

func main() {
	var specFile string
	flag.StringVar(&specFile, "f", "", "Path to the specific file")
	flag.Parse()
	if specFile == "" {
		fmt.Println("spec file should be provided to generate the api")
	}
	byt, err := ioutil.ReadFile(specFile)
	if err != nil {
		fmt.Printf("Spec file could not be read %s", err.Error())
	}
	var spec models.Spec
	if err := json.Unmarshal(byt, &spec); err != nil {
		panic(err)
	}
	modelList := specparser.ParseRefObjectMap(spec.Refs)
	fmt.Printf("The spec file is %+v", modelList)
	//creator.ModelCreator(modelMap, "test")

}

type SomeSuccessResponse struct {
	Sometrribute1 []interface{}
	Sometrribute2 string
	Sometrribute3 float64
}
