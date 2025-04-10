package configs

import (
	"log"
	"os"
	"time"
)

type DBConfig struct {
	PostgresDSN     string
	MaxConns        int32
	MinConns        int32
	MaxConnIdleTime time.Duration
}

var DB DBConfig

func loadDBConfig() {
	DB = DBConfig{
		PostgresDSN:     os.Getenv("POSTGRES_DSN"),
		MaxConns:        10,
		MinConns:        2,
		MaxConnIdleTime: 30 * time.Minute,
	}
	if DB.PostgresDSN == "" {
		log.Fatal("POSTGRES_DSN is not set")
	}

	log.Println("Database config loaded")
}
