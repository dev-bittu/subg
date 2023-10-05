package config

import (
  "encoding/json"
  _ "embed"
)

//go:embed config.json
var cfg []byte

var Config map[string]interface{}

func init() {
  err := json.Unmarshal(cfg, &Config)
  if err != nil {
    println("Cant load config")
    panic(err)
  }
}
