package config

import "testing"

type yamlConfigLoad struct {
	Host  string 	`yaml:"Host"`
}

type jsonConfigLoad struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type tomlConfigLoad struct {
	Server string  		`toml:"server"`
	Port string			`toml:"port"`
	Database string		`toml:"database"`
	User string			`toml:"user"`
	Password string		`toml:"Password"`
}

func TestLoadConfig(t *testing.T) {
	err := LoadConfig("test_configs","app.yaml", &yamlConfigLoad{}, true)
	if err !=nil{
		t.Errorf("Error loading YAML config, load function test [ERROR]: %s",err)
	}

	err = LoadConfig("test_configs","app.json", &jsonConfigLoad{}, true)
	if err !=nil{
		t.Errorf("Error loading JSON config, load function test [ERROR]: %s",err)
	}

	err = LoadConfig("test_configs","app.toml", &tomlConfigLoad{}, true)
	if err !=nil{
		t.Errorf("Error loading TOML config, load function test [ERROR]: %s",err)
	}
}

func BenchmarkLoadConfig(b *testing.B) {
	for i:=0; i<b.N; i++ {
		err := LoadConfig("test_configs","app.yaml", &yamlConfigLoad{}, true)
		if err !=nil{
			b.Errorf("Error loading YAML config, load function test [ERROR]: %s",err)
		}

		err = LoadConfig("test_configs","app.json", &jsonConfigLoad{}, true)
		if err !=nil{
			b.Errorf("Error loading JSON config, load function test [ERROR]: %s",err)
		}

		err = LoadConfig("test_configs","app.toml", &tomlConfigLoad{}, true)
		if err !=nil{
			b.Errorf("Error loading TOML config, load function test [ERROR]: %s",err)
		}
		if err != nil {
			b.Error("Failed: LoadConfig Function [ERROR] :", err)
		}
	}

}

func TestLoad(t *testing.T){
	err := load("test_configs","app.yaml", &yamlConfigLoad{}, true)
	if err !=nil{
		t.Errorf("Error loading YAML config, load function test [ERROR]: %s",err)
	}

	err = load("test_configs","app.json", &jsonConfigLoad{}, true)
	if err !=nil{
		t.Errorf("Error loading JSON config, load function test [ERROR]: %s",err)
	}

	err = load("test_configs","app.toml", &tomlConfigLoad{}, true)
	if err !=nil{
		t.Errorf("Error loading TOML config, load function test [ERROR]: %s",err)
	}
}

func BenchmarkLoad(b *testing.B){
	for i:=0; i<b.N; i++ {
		err := load("test_configs","app.yaml", &yamlConfigLoad{}, true)
		if err !=nil{
			b.Errorf("Error loading YAML config, load function test [ERROR]: %s",err)
		}

		err = load("test_configs","app.json", &jsonConfigLoad{}, true)
		if err !=nil{
			b.Errorf("Error loading JSON config, load function test [ERROR]: %s",err)
		}

		err = load("test_configs","app.toml", &tomlConfigLoad{}, true)
		if err !=nil{
			b.Errorf("Error loading TOML config, load function test [ERROR]: %s",err)
		}
		if err != nil {
			b.Error("Failed: load Function [ERROR] :", err)
		}
	}
}