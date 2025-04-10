package configs

import (
	"log"
	"os"
)

type AppConfig struct {
	URL      string
	GRPCPort string
}

var App AppConfig

func loadAppConfig() {
	App = AppConfig{
		URL:      os.Getenv("APP_URL"),
		GRPCPort: os.Getenv("GRPC_PORT"),
	}

	if App.URL == "" {
		log.Fatal("APP_URL is not set")
	}
	if App.GRPCPort == "" {
		log.Fatal("GRPC_PORT is not set")
	}

	log.Println("App config loaded")
}
