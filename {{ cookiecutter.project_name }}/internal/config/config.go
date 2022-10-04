package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var once sync.Once
var data = make(map[interface{}]interface{})

func Config() map[interface{}]interface{} {
	once.Do(func()  {
		path := os.Getenv("CONFIG_FILE")
		if path == "" {
			path = "/etc/{{ cookiecutter.project_name }}/config.yaml"
		}
		yfile, err := os.ReadFile(path)
	
		if err != nil {
			log.Fatal(err)
		}
	
		err2 := yaml.Unmarshal(yfile, &data)
		if err2 != nil {
			log.Fatal(err)
		}			
	})

	return data
}
