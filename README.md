# Config
Friendly GO YAML Configuration loader.

## Installation
````
go get github.com/anjulapaulus/config
````

## Usage

Folder Structure

````
config/
    config.yaml
main.go
````

Implementation
````
import (
	"fmt"
	"github.com/anjulapaulus/config"
)

type AppConfig struct {
	Host     string  `yaml:"host"`

}

func main(){
	r,err := config.LoadYAMLConfig("config","config.yaml", AppConfig{})
	if err != nil{
		panic("Could not load configuration")
	}
	fmt.Println(r)
}
````

