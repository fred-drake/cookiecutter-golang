package main

import (
	"fmt"
	"os"

	"github.com/{{ cookiecutter.github_name }}/{{ cookiecutter.project_name }}/internal/config"
	"github.com/{{ cookiecutter.github_name }}/{{ cookiecutter.project_name }}/internal/log"
)

func main() {
	logger := log.GetLogger()
	logger.Info("Starting.")
	validate()

	var config = config.Config()
	logger.Info(config["foo"])

}

func validate() {
	logger := log.GetLogger()
	requiredEnvs := []string{"REQUIRED", "VARIABLES", "HERE"}
	for _, r := range requiredEnvs {
		if os.Getenv(r) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s must be defined!", r))
			os.Exit(1)
		}
	}
}
