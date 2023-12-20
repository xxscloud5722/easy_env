package app

import (
	"embed"
	_ "embed"
	"gopkg.in/yaml.v3"
	"log"
)

//go:embed app.yaml
var appConfig string

//go:embed 404.html
var template404 string

//go:embed 500.html
var template500 string

//go:embed files.html
var templateFiles string

//go:embed admin/*
var AdminDir embed.FS

var (
	app      *App
	template map[string]string
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

	template = map[string]string{}
	template["404"] = template404
	template["500"] = template500
	template["files"] = templateFiles
}

func Info() *App {
	return app
}

func Template() map[string]string {
	return template
}
