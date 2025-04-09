package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	URL string
}

type DBConfig struct {
	PostgresDSN     string
	MaxConns        int
	MinConns        int
	MaxConnIdleTime int
}

var App AppConfig
var DB DBConfig

// LoadConfig loads the application configuration from environment variables.
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file not found")
	} else {
		loadAppConfig()
		loadDBConfig()
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

func loadDBConfig() {
	DB = DBConfig{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
	}
	if DB.PostgresDSN == "" {
		log.Fatal("POSTGRES_DSN is not set")
	}
	log.Println("Database config loaded")
}
