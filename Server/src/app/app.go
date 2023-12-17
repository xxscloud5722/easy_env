package app

import (
	_ "embed"
	"gopkg.in/yaml.v3"
	"log"
)

//go:embed app.yaml
var appConfig string

var (
	app *App
)

type App struct {
	Version string `yaml:"version"`
}

func init() {
	var config App
	err := yaml.Unmarshal([]byte(appConfig), &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}
	app = &config
}

func Info() *App {
	return app
}
