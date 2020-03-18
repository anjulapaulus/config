package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"strings"
)

func readFile(folderName string, fileName string, config interface{}, checkUnmatched bool) (err error) {
	filePath := fmt.Sprintf("./" + folderName + "/" + fileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot Read config file [ERROR]: %s",err))
	}
	switch {
	case strings.HasSuffix(fileName, ".yaml") || strings.HasSuffix(fileName, ".yml"):
		if checkUnmatched {
			return yaml.UnmarshalStrict(content, config)
		}
		return yaml.Unmarshal(content, config)
	case strings.HasSuffix(fileName, ".toml"):
		return unmarshalToml(content, config, checkUnmatched)
	case strings.HasSuffix(fileName, ".json"):
		return unmarshalJSON(content, config, checkUnmatched)
	default:
		return errors.New("unable to decode file type")
	}

	return nil

}

func unmarshalJSON(content []byte, config interface{}, checkUnmatched bool) error {
	reader := strings.NewReader(string(content))
	decoder := json.NewDecoder(reader)

	if checkUnmatched {
		decoder.DisallowUnknownFields()
	}

	err := decoder.Decode(config)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

func unmarshalToml(content []byte, config interface{}, checkUnmatched bool) error {
	metadata, err := toml.Decode(string(content), config)
	if err == nil && len(metadata.Undecoded()) > 0 && checkUnmatched {
		return errors.New("unable to unmarshal toml")
	}
	return nil
}
