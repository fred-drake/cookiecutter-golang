module github.com/{{ cookiecutter.github_name }}/{{ cookiecutter.project_name }}

go {{ cookiecutter.go_version }}

require (
	github.com/gookit/config/v2 v2.1.6
	github.com/sirupsen/logrus v1.9.0
)

require (
	github.com/gookit/goutil v0.5.12 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
