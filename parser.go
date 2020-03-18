package config

import (
	"fmt"
	"reflect"
)

func LoadConfig(folderName string,fileName string, config interface{}, checkUnmatchedKeys bool) (err error) {
	return load(folderName,fileName, config, checkUnmatchedKeys)
}

func load(folderName string,fileName string, config interface{},checkUnmatchedKeys bool)(err error){
	value := reflect.Indirect(reflect.ValueOf(config))
	if !value.CanAddr() {
		return fmt.Errorf("config should be addressable")
	}
	err = readFile(folderName,fileName, config, checkUnmatchedKeys)
	if err != nil{
		return err
	}
	return nil
}

