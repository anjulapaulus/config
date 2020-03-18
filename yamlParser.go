package config

import (
	"gopkg.in/yaml.v2"
)

func LoadYAMLConfig(folderName string,fileName string, config interface{}) (interface{}, error) {

	content, err := read(folderName,fileName)
	if err != nil{
		return nil, err
	}

	cfg := config

	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
