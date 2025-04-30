package configs

import (
	"log"
	"os"
)

type DatabaseConfig struct {
	PostgresDSN PostgresDSN
}

type PostgresDSN struct {
	Product string
	Score   string
}

var Database DatabaseConfig

func loadDatabaseConfig() {
	PostgresProductDSN := os.Getenv("POSTGRES_PRODUCT_DSN")
	PostgresScoreDSN := os.Getenv("POSTGRES_SCORE_DSN")
	if PostgresProductDSN == "" || PostgresScoreDSN == "" {
		log.Fatal("POSTGRES_PRODUCT_DSN or POSTGRES_SCORE_DSN not set")
	}

	Database.PostgresDSN = PostgresDSN{
		Product: PostgresProductDSN,
		Score:   PostgresScoreDSN,
	}

	log.Println("Database config loaded")
}
