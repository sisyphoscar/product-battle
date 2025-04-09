package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	URL string
}

var App AppConfig

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file not found")
	} else {
		loadAppConfig()
	}
}

func loadAppConfig() {
	App = AppConfig{
		URL: os.Getenv("APP_URL"),
	}
	if App.URL == "" {
		log.Fatal("APP_URL is not set")
	}
	log.Println("App config loaded")
}
