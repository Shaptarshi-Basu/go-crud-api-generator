package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"go-crud-api-generator/models"
	"encoding/json"
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
	fmt.Printf("The spec file is %+v", spec)
}
