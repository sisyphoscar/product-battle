package configs

import (
	"log"
	"os"
)

type AppConfig struct {
	Port     string
	GRPCPort string
}

var App AppConfig

func loadAppConfig() {
	App = AppConfig{
		Port:     os.Getenv("APP_PORT"),
		GRPCPort: os.Getenv("GRPC_PORT"),
	}

	if App.Port == "" {
		log.Fatal("APP_PORT is not set")
	}
	if App.GRPCPort == "" {
		log.Fatal("GRPC_PORT is not set")
	}

	log.Println("App config loaded")
}
