[![Build Status](https://travis-ci.org/anjulapaulus/config.svg?branch=master)](https://travis-ci.org/anjulapaulus/config)
[![HitCount](http://hits.dwyl.com/anjulapaulus/config.svg)](http://hits.dwyl.com/anjulapaulus/config)
# Config
Friendly GO YAML, JSON and TOML Configuration loader.

## Installation
````
go get github.com/anjulapaulus/config
````

## Usage

Folder Structure

````
config/
    config.yaml
    config.json
    config.toml
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

func main() {
    #if you want to strictly check on unmatched keys in config file
    #else false
	config.LoadConfig("config","config.yaml",&Config,true)
	fmt.Println(Cfg.APPName)
}
````

## Benchmarks
go test -bench=Read reader_test.go parser.go reader.go parser_test.go
````
goos: darwin
goarch: amd64
BenchmarkReadFile-8                10000            113305 ns/op
PASS
ok      command-line-arguments  1.308s
````
go test -bench=Load reader_test.go parser.go reader.go parser_test.go
````
goos: darwin
goarch: amd64
BenchmarkLoadConfig-8              10000            111349 ns/op
BenchmarkLoad-8                    10000            110499 ns/op
PASS
ok      command-line-arguments  2.382s
````

