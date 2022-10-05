package main

import (
	"github.com/{{ cookiecutter.github_name }}/{{ cookiecutter.project_name }}/internal/initialization"
	"github.com/{{ cookiecutter.github_name }}/{{ cookiecutter.project_name }}/internal/logger"
)

func main() {
	// Set the list of command line flags (with optional type separated by colon)
	cliFlags := []string{
		"configFile",
		"logFormat",
		"logLevel",
	}

	// Set the mapping of environment variables to configuration keys.
	envToKeys := map[string]string{
		"CONFIG_FILE": "configFile",
		"LOG_FORMAT": "logFormat",
		"LOG_LEVEL": "logLevel",
	}

	// Set hard-coded defaults if not overridden anywhere else
	configDefaults := map[string]interface{}{
		"configFile": "/etc/my-project/config.yaml",
		"logFormat": "text",
		"logLevel": "info",
	}

	// Required config keys.  If any are defined here but not set in the configuration, 
	// the application will panic.
	requiredConfigKeys := []string{}

	initialization.Configuration(cliFlags, envToKeys, configDefaults)

	log := logger.GetLogger()
	log.Info("Starting.")
	initialization.Validate(requiredConfigKeys)
}
