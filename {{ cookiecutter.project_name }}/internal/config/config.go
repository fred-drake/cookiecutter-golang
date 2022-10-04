package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v3"
)

var once sync.Once
var data = make(map[interface{}]interface{})

func Config() map[interface{}]interface{} {
	once.Do(func ()  {
		yfile, err := ioutil.ReadFile("config.yaml")
	
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
