package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-crud-api-generator/creator"
	"go-crud-api-generator/models"
	specparser "go-crud-api-generator/parser"
	"io/ioutil"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Print(router)
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
	apiInfos := specparser.ApiInfo{}
	modelList := specparser.ParseRefObjectMap(spec.Refs)
	fmt.Printf("The spec file is %v", modelList)
	creator.ModelCreator(modelList, "test")
	mainFuncStr := apiInfos.CreateMainFunc(spec.Paths)
	fmt.Printf("Main is %v", mainFuncStr)
	creator.MainMethodCreator(mainFuncStr, "test")
	handlerFuncStr := apiInfos.CreateHandlerFunc()
	creator.HandlerCreator(handlerFuncStr, "test")
}
