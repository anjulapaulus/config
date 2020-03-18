package config

import (
	"fmt"
	"io/ioutil"
)

func read(folderName string, fileName string) ([]byte, error) {
	filePath := fmt.Sprintf("./"+folderName+"/"+fileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content,nil

}
