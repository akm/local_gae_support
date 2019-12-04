package localgaesupport

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type AppYaml struct {
	Runtime  string          `yaml:"runtime"`
	Service  string          `yaml:"service"`
	Main     string          `yaml:"main"`
	Handlers AppYamlHandlers `yaml:"handlers"`
}

func ParseAppYaml(path string) (*AppYaml, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read YAML file because of %v", err)
		return nil, err
	}
	var r AppYaml
	if err := yaml.Unmarshal(b, &r); err != nil {
		log.Printf("Failed to unmarshal YAML file because of %v", err)
		return nil, err
	}

	return &r, nil
}
