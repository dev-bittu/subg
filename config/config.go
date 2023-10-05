package config

import (
	_ "embed"
	"encoding/json"
)

//go:embed config.json
var cfg []byte

var Config map[string]interface{}

func init() {
	// load config & assign to Config
	err := json.Unmarshal(cfg, &Config)
	if err != nil {
		println("Cant load config")
		panic(err)
	}
}
