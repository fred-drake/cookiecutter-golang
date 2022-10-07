package main

import (
	"github.com/fred-drake/{{ cookiecutter.project_name }}/internal/initialization"
	"github.com/fred-drake/{{ cookiecutter.project_name }}/internal/logger"
)

func main() {
	// Set the list of command line flags (with optional type separated by colon)
	configKeys := []string{
		"configFile",
		"logFormat",
		"logLevel",
		"namespace",
		"pgHost",
		"pgDatabase",
		"pgPassword",
		"pgBackupFile",
		"mysqlUser",
		"mysqlPassword",
		"mysqlHost",
		"backupDirectory",
		"sleepBetweenBackups:int",
		"pgPort:int",
	}

	// Set hard-coded defaults if not overridden anywhere else
	configDefaults := map[string]interface{}{
		"configFile": "/etc/{{ cookiecutter.project_name }}/config.yaml",
		"logFormat":  "text",
		"logLevel":   "info",
	}

	// Required config keys.  If any are defined here but not set in the configuration,
	// the application will panic.
	requiredConfigKeys := []string{"namespace"}

	var settings = &initialization.ConfigurationSettings{
		ConfigKeys:    configKeys,
		DefaultValues: configDefaults,
	}
	initialization.Configuration(settings)

	log := logger.GetLogger()
	log.Info("Starting.")
	initialization.Validate(requiredConfigKeys)
}
