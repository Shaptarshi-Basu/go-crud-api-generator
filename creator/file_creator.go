package creator

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ModelCreator(modelList []string, project string) {
	err := directoryCreator(fmt.Sprintf("%s/model", project))
	if err != nil {
		fmt.Printf(err.Error())
	}
	f, err := os.OpenFile(fmt.Sprintf("%s/model/models.go", project),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s\n", "package model")); err != nil {
		log.Println(err)
	}
	for _, v := range modelList {
		if _, err := f.WriteString(fmt.Sprintf("%s\n", v)); err != nil {
			log.Println(err)
		}
	}

}

func MainMethodCreator(mainFunc, project, port string) {
	mainFunc = strings.ReplaceAll(mainFunc, "<PROJECT_NAME>", project)
	mainFunc = strings.ReplaceAll(mainFunc, "<PORT>", port)
	f, err := os.OpenFile(fmt.Sprintf("%s/main.go", project),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s\n", mainFunc)); err != nil {
		log.Println(err)
	}

}

func HandlerCreator(mainFunc, project string) {
	err := directoryCreator(fmt.Sprintf("%s/handler", project))
	if err != nil {
		fmt.Printf(err.Error())
	}
	f, err := os.OpenFile(fmt.Sprintf("%s/handler/handler.go", project),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s\n", mainFunc)); err != nil {
		log.Println(err)
	}

}
