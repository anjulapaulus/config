package config

import (
	"io/ioutil"
)

func read(path string, file string) ([]byte, error) {
	content, err := ioutil.ReadFile(path + file)
	if err != nil {
		return nil, err
	}
	return content,nil

}
