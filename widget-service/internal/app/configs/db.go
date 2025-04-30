package configs

import (
	"log"
	"os"
	"time"
)

type DBConfig struct {
	PostgresDSN     PostgresDSN
	MaxConns        int32
	MinConns        int32
	MaxConnIdleTime time.Duration
}

type PostgresDSN struct {
	Product string
	Score   string
}

var DB DBConfig

func loadDBConfig() {
	PostgresProductDSN := os.Getenv("POSTGRES_PRODUCT_DSN")
	PostgresScoreDSN := os.Getenv("POSTGRES_SCORE_DSN")
	if PostgresProductDSN == "" || PostgresScoreDSN == "" {
		log.Fatal("POSTGRES_PRODUCT_DSN or POSTGRES_SCORE_DSN not set")
	}

	DB.PostgresDSN = PostgresDSN{
		Product: PostgresProductDSN,
		Score:   PostgresScoreDSN,
	}

	DB.MaxConns = 10
	DB.MinConns = 2
	DB.MaxConnIdleTime = 30 * time.Second

	log.Println("Database config loaded")
}
