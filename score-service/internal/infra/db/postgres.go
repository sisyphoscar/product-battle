package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sisyphoscar/product-battle/score-service/internal/configs"
)

const (
	MAX_RETRIES        = 5
	RETRY_INTERVAL     = 5 * time.Second
	CONNECTION_TIMEOUT = 10 * time.Second
)

// NewPostgres initializes a new PostgreSQL connection pool with retry logic.
func NewPostgres() (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error

	config, err := prepareConfig()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), CONNECTION_TIMEOUT)
	defer cancel()

	pool, err = connectPostgresWithRetry(ctx, config)
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	return pool, nil
}

// prepareConfig prepares the PostgreSQL configuration from the environment variables.
func prepareConfig() (*pgxpool.Config, error) {
	config, parseErr := pgxpool.ParseConfig(configs.DB.PostgresDSN)
	if parseErr != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", parseErr)
	}

	config.MaxConns = configs.DB.MaxConns
	config.MinConns = configs.DB.MinConns
	config.MaxConnIdleTime = configs.DB.MaxConnIdleTime

	return config, nil
}

// connectPostgresWithRetry attempts to connect to PostgreSQL with retry logic.
func connectPostgresWithRetry(ctx context.Context, config *pgxpool.Config) (*pgxpool.Pool, error) {
	for i := 0; i < MAX_RETRIES; i++ {
		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err == nil {
			err = pool.Ping(ctx)
			if err == nil {
				log.Println("Database connected")
				return pool, nil
			}
		}

		log.Printf("Retrying to connect to database (%d/%d): %v\n", i+1, MAX_RETRIES, err)
		time.Sleep(RETRY_INTERVAL)
	}

	return nil, errors.New("database connection failed")
}
