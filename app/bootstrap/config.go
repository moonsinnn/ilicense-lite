package bootstrap

import (
	"io/ioutil"
	"log"

	"ilicense-lite/config"

	"gopkg.in/yaml.v3"
)

func InitConfig(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	err = yaml.Unmarshal(data, &config.Config)
	if err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}
}
