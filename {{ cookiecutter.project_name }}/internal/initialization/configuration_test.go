package initialization

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gookit/config/v2"
	"github.com/stretchr/testify/assert"
)

func TestConfigurationEnvPriority(t *testing.T) {
	defer config.ClearAll()

	file := "/tmp/" + uuid.New().String() + ".yaml"
	settings := &ConfigurationSettings{
		ConfigKeys: []string{"onomatopoeia"},
		DefaultValues: map[string]interface{}{
			"configFile":   file,
			"onomatopoeia": "pow",
		},
	}
	os.Clearenv()
	os.Setenv("ONOMATOPOEIA", "swoosh")
	yaml := "onomatopoeia: bang"
	f, err := os.Create(file)
	if err != nil {
		assert.Fail(t, "Couldn't create file "+file)
		return
	}
	defer os.Remove(file)
	f.WriteString(yaml + "\n")
	f.Close()

	Configuration(settings)
	assert.Equal(t, "swoosh", config.String("onomatopoeia"))
}

func TestConfigurationFilePriority(t *testing.T) {
	defer config.ClearAll()

	file := "/tmp/" + uuid.New().String() + ".yaml"
	settings := &ConfigurationSettings{
		ConfigKeys: []string{},
		DefaultValues: map[string]interface{}{
			"configFile":   file,
			"onomatopoeia": "pow",
		},
	}
	os.Clearenv()
	yaml := "onomatopoeia: bang"
	f, err := os.Create(file)
	if err != nil {
		assert.Fail(t, "Couldn't create file "+file)
		return
	}
	defer os.Remove(file)
	f.WriteString(yaml + "\n")
	f.Close()

	Configuration(settings)
	assert.Equal(t, "bang", config.String("onomatopoeia"))
}

func TestConfigurationDefaultPriority(t *testing.T) {
	defer config.ClearAll()

	settings := &ConfigurationSettings{
		ConfigKeys: []string{},
		DefaultValues: map[string]interface{}{
			"onomatopoeia": "pow",
		},
	}
	os.Clearenv()

	Configuration(settings)
	assert.Equal(t, "pow", config.String("onomatopoeia"))
}

func TestEnvKeysFromCLIOptions(t *testing.T) {
	result := envKeysFromConfigKeys([]string{"foo", "fooBar", "FooBar", "FooBarBaz"})
	expected := map[string]string{
		"foo":       "FOO",
		"fooBar":    "FOO_BAR",
		"FooBar":    "FOO_BAR",
		"FooBarBaz": "FOO_BAR_BAZ",
	}

	assert.Equal(t, expected, result)
}
