package config

import (
	"gopkg.in/yaml.v2"
)

func LoadYAMLConfig(folderPath string, file string, config interface{}) (interface{}, error) {

	content := read(folderPath, file)

	cfg := config

	err := yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
