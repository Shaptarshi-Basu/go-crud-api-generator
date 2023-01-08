package main

import (
	"fmt"
	specgenerator "go-crud-api-generator/generator"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No filepath for env variables")
		return
	}
	filepath := os.Args[1]
	err := godotenv.Load(filepath)
	if err != nil {
		log.Fatal(".env file must be present")
	}
	err = specgenerator.TriggerApiGenerator()
	if err != nil {
		fmt.Print("Error returned ", err.Error())
	}
}
