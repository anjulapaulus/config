package config

import (
	"testing"
)


type yamlConfig struct {
	Host  string 	`yaml:"Host"`
}

type jsonConfig struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type tomlConfig struct {
	Server string  		`toml:"server"`
	Port string			`toml:"port"`
	Database string		`toml:"database"`
	User string			`toml:"user"`
	Password string		`toml:"Password"`
}

func TestReadFile(t *testing.T){
	err := readFile("test_configs","app.yaml", &yamlConfig{}, true)
	if err !=nil{
		t.Errorf("Error loading YAML config [ERROR]: %s",err)
	}

	err = readFile("test_configs","app.json", &jsonConfig{}, true)
	if err !=nil{
		t.Errorf("Error loading JSON config [ERROR]: %s",err)
	}

	err = readFile("test_configs","app.toml", &tomlConfig{}, true)
	if err !=nil{
		t.Errorf("Error loading TOML config [ERROR]: %s",err)
	}

}

func BenchmarkReadFile(b *testing.B){
	for i:=0; i<b.N; i++ {
		err := readFile("test_configs","app.yaml", &yamlConfigLoad{}, true)
		if err !=nil{
			b.Errorf("Error loading YAML config, load function test [ERROR]: %s",err)
		}

		err = readFile("test_configs","app.json", &jsonConfigLoad{}, true)
		if err !=nil{
			b.Errorf("Error loading JSON config, load function test [ERROR]: %s",err)
		}

		err = readFile("test_configs","app.toml", &tomlConfigLoad{}, true)
		if err !=nil{
			b.Errorf("Error loading TOML config, load function test [ERROR]: %s",err)
		}
		if err != nil {
			b.Error("Failed: LoadConfig Function [ERROR] :", err)
		}
	}
}
