package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

func readFile(folderName string, fileName string, config interface{}, checkUnmatched bool) (err error) {
	filePath := fmt.Sprintf("./"+folderName+"/"+fileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	switch {
	case strings.HasSuffix(fileName, ".yaml") || strings.HasSuffix(fileName, ".yml"):
		if checkUnmatched{
			return yaml.UnmarshalStrict(content,config)
		}
		return yaml.Unmarshal(content, config)

	}
	return nil

}

