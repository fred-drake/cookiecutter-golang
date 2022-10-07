package initialization

import (
	"os"
	"testing"

	"github.com/gookit/config/v2"
	"github.com/stretchr/testify/assert"
)

func TestValidateWithEmptyConfigSettings(t *testing.T) {
	defer config.ClearAll()
	os.Clearenv()
	Configuration(NewSettings())
	assert.True(t, Validate([]string{}))
}

func TestNoConfigVariables(t *testing.T) {
	defer config.ClearAll()
	os.Clearenv()
	settings := &ConfigurationSettings{
		ConfigKeys:    []string{"thisIsRequired"},
		DefaultValues: map[string]interface{}{},
	}
	Configuration(settings)
	assert.False(t, Validate([]string{"thisIsRequired"}))
}

func TestValidConfigVariables(t *testing.T) {
	defer config.ClearAll()
	os.Clearenv()
	settings := &ConfigurationSettings{
		ConfigKeys:    []string{"thisIsRequired"},
		DefaultValues: map[string]interface{}{},
	}
	os.Setenv("THIS_IS_REQUIRED", "party")
	Configuration(settings)
	assert.True(t, Validate([]string{"thisIsRequired"}))
}
