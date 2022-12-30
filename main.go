package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var specFile string
	flag.StringVar(&specFile, "f", "", "Path to the specific file")
	flag.Parse()
	fmt.Print("Spec file is ", specFile)
	spec, err := ioutil.ReadFile(specFile)
	if err != nil {
		fmt.Printf("Spec file could not be read %s", err.Error())
	}
	fmt.Println(string(spec))
}
