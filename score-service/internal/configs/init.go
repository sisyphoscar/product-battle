package configs

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig loads the configuration from the environment variables
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file not found")
	}

	loadAppConfig()
	loadQueueConfig()
	loadDatabaseConfig()
}
