package configs

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig loads the application configuration from environment variables.
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file not found")
	}

	loadAppConfig()
	loadDBConfig()
}
