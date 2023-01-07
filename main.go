package main

import (
	"fmt"
	specgenerator "go-crud-api-generator/generator"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file must be present")
	}
	err = specgenerator.TriggerApiGenerator()
	if err != nil {
		fmt.Print("Error returned ", err.Error())
	}
}
