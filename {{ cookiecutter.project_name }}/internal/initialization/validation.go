package initialization

import (
	"fmt"

	"github.com/fred-drake/{{ cookiecutter.project_name }}/internal/logger"
	"github.com/gookit/config/v2"
)

func Validate(requiredConfigKeys []string) bool {
	log := logger.GetLogger()

	for _, r := range requiredConfigKeys {
		if !config.Exists(r, false) {
			err := fmt.Sprintf("Config variable %s must be defined!", r)
			log.Error(fmt.Sprintf(err, r))

			return false
		}
	}

	return true
}
