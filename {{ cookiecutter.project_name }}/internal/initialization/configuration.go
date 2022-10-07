package initialization

import (
	"os"
	"strings"
	"sync"
	"unicode"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

type ConfigurationSettings struct {
	ConfigKeys    []string
	DefaultValues map[string]interface{}
}

var once sync.Once

func Configuration(configValues *ConfigurationSettings) {
	// We need to get the "configFile" key in order to pull from
	// the YAML file.  Pull the configs from other sources, then
	// re-apply them so they will follow the correct overriding order.

	// Config will panic if CLI is read twice, so store it in a JSON
	// for subsequent calls.
	loadCLIKeys(configValues.ConfigKeys)
	cliJson := config.Default().ToJSON()

	// Check default, env and cli sources to get configFile variable
	config.ClearAll()
	for k, v := range configValues.DefaultValues {
		config.Set(k, v)
	}
	// config.SetData(configValues.DefaultValues)
	loadEnvKeys(configValues.ConfigKeys)
	config.LoadSources(config.JSON, []byte(cliJson))

	loadYamlFile()

	// Re-apply ENV and CLI to ensure any conflicts
	// overwrite the config file's values
	loadEnvKeys(configValues.ConfigKeys)
	config.LoadSources(config.JSON, []byte(cliJson))
}

func envKeysFromConfigKeys(configKeys []string) map[string]string {
	configToEnvKeys := map[string]string{}

	for _, configKey := range configKeys {
		envKey := ""
		for i, c := range configKey {
			if unicode.IsUpper(c) && i > 0 {
				envKey = envKey + "_" + string(c)
			} else {
				envKey = envKey + strings.ToUpper(string(c))
			}
		}
		configToEnvKeys[configKey] = envKey
	}

	return configToEnvKeys
}

func loadEnvKeys(configKeys []string) {
	// Override with environment variables
	envKeys := envKeysFromConfigKeys(configKeys)
	for configKey, envKey := range envKeys {
		if os.Getenv(envKey) != "" {
			config.Set(configKey, os.Getenv(envKey))
		}
	}
}

func loadCLIKeys(configKeys []string) {
	// If config.LoadFlags() is called multiple times, the config library
	// will panic.  This is only intended to be called once, but it does get
	// called multiple times during several unit tests.  We don't make CLI
	// calls during tests so wrap this in once.Do() to keep it from failing.
	// The alternative is to always ignore err, which we do not want to do.
	once.Do(func() {
		// Override with command line flags
		err := config.LoadFlags(configKeys)
		if err != nil {
			panic(err)
		}
	})
}

func loadYamlFile() {
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
