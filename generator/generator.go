package generator

import (
	"fmt"
	"go-crud-api-generator/creator"
	"go-crud-api-generator/models"
	"os"
)

type envConfig struct {
	SpecFilePath string
	Port         string
	ProjectName  string
	GoVersion    string
}
type apiGenrerator struct {
	spec models.Spec
	env  envConfig
}

func TriggerApiGenerator() error {
	env := envConfig{}
	err := env.populateAPIInfosFromEnv()
	if err != nil {
		return err
	}
	spec := models.Spec{}
	err = spec.FetchDetailsFromSpecFile(env.SpecFilePath)
	if err != nil {
		return err
	}
	apiGen := apiGenrerator{spec: spec, env: env}
	apiGen.orchestrateCodeGeneration()
	return nil
}

func (ag *apiGenrerator) orchestrateCodeGeneration() {
	if len(ag.spec.Refs) != 0 {
		modeList := generateModelStructs(ag.spec.Refs)
		if len(modeList) != 0 {
			creator.ModelCreator(modeList, ag.env.ProjectName)
		}
	}
	creator.HandlerCreator(ag.generateHandlers(), ag.env.ProjectName)

	creator.MainMethodCreator(ag.generateMainFunc(), ag.env.ProjectName, ag.env.Port)

	creator.GoModCreator(ag.generateModFile(), ag.env.ProjectName)

}

func (env *envConfig) populateAPIInfosFromEnv() error {
	specFilePath := os.Getenv("SPEC_FILE_PATH")
	if specFilePath == "" {
		return fmt.Errorf("Spec file path cannot be empty for api generation")
	} else {
		env.SpecFilePath = specFilePath
	}
	port := os.Getenv("PORT")
	if port == "" {
		env.Port = "8080"
	} else {
		env.Port = port
	}
	projectName := os.Getenv("PROJECT_NAME")
	if projectName == "" {
		return fmt.Errorf("The generated project name cannot be empty so has to be mentioned")
	} else {
		env.ProjectName = projectName
	}

	goVersion := os.Getenv("GO_VERSION")
	if goVersion == "" {
		env.GoVersion = "1.13"
	} else {
		env.GoVersion = goVersion
	}
	return nil
}
