package configs

import (
	"log"
	"os"
)

type DatabaseConfig struct {
	PostgresDSN string
}

var Database DatabaseConfig

func loadDatabaseConfig() {
	Database = DatabaseConfig{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
	}

	if Database.PostgresDSN == "" {
		log.Fatal("POSTGRES_DSN is not set")
	}

	log.Println("Database config loaded")
}
