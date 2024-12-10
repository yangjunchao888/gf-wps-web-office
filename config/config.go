package config

import (
	"os"
)

var (
	configFile = ""
)

func Init() {
	mode := os.Getenv("MODE")
	switch mode {
	case "test":
		configFile = "config-test.yaml"
	case "prod":
		configFile = "config-prod.yaml"
	case "local":
		configFile = "config-local.yaml"
	default:
		configFile = "config-test.yaml"
	}
}
