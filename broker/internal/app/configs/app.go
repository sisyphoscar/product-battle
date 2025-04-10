package configs

import (
	"log"
	"os"
)

type AppConfig struct {
	URL string
}

var App AppConfig

func loadAppConfig() {
	App = AppConfig{
		URL: os.Getenv("APP_URL"),
	}
	if App.URL == "" {
		log.Fatal("APP_URL is not set")
	}

	log.Println("App config loaded")
}
