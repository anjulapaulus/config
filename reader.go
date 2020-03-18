package config

import (
	"io/ioutil"
)

func read(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content,nil

}
