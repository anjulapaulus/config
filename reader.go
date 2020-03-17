package config

import (
	"fmt"
	"io/ioutil"
)

func read(path string, file string) []byte {
	folder := fmt.Sprintf("./" + path + "/")
	content, err := ioutil.ReadFile(folder + file)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return content

}
