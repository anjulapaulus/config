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

import (
	"fmt"
	"github.com/anjulapaulus/config"
)

type Config struct {
	APPName string `yaml:"appname"`

}

var Cfg Config

func main() {
	config.LoadConfig("config","config.yaml",&Cfg,true)
	fmt.Println(Cfg.APPName)
}
````

