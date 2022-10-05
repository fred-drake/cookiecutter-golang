package initialization

import (
	"fmt"
	"os"

	"github.com/{{ cookiecutter.github_name }}/{{ cookiecutter.project_name }}/internal/logger"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

func Configuration(cliFlags []string, envKeys map[string]string, defaults map[string]interface{}) {
	// Can only read CLI once so store it in a JSON to pull later
	configCLI(cliFlags)
	cliJson := config.Default().ToJSON()
	config.ClearAll()

	// Check sources to get configFile variable
	config.SetData(defaults) // Set hard-coded defaults
	configEnv(envKeys) // Set environment variables
	config.LoadSources(config.JSON, []byte(cliJson)) // Load CLI
	configYamlFile() // Pull via config file, if applicable

	// Re-apply ENV and CLI to ensure any conflicts
	// overwrite the config file
	configEnv(envKeys) // Set environment variables
	config.LoadSources(config.JSON, []byte(cliJson)) // Load CLI
}

func Validate(requiredConfigKeys []string) {
	log := logger.GetLogger()
	for _, r := range requiredConfigKeys {
		if !config.Exists(r, false) {
			err := fmt.Sprintf("Config variable %s must be defined!", r)
			log.Fatal(fmt.Sprintf(err, r))
			panic(err)
		}
	}
}

func configEnv(envKeys map[string]string) {
	// Override with environment variables
	for k, v := range envKeys {
		if os.Getenv(k) != "" {
			config.Set(v, os.Getenv(k))
		}
	}
}

func configCLI(cliFlags []string) {
	// Override with command line flags
	err := config.LoadFlags(cliFlags)
	if err != nil {
		panic(err)
	}
}

func configYamlFile() {
	// Override with config file
	path := config.String("configFile")
	_, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err == nil {
		config.AddDriver(yamlv3.Driver)
		err := config.LoadFiles(path)
		if err != nil {
			panic(err)
		}
	}
}