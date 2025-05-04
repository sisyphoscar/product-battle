package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sisyphoscar/product-battle/bi-service/internal/app/configs"
)

const (
	MAX_RETRIES        = 5
	RETRY_INTERVAL     = 5 * time.Second
	CONNECTION_TIMEOUT = 10 * time.Second
)

// NewProductPostgres creates a new PostgreSQL connection pool for the product database
func NewProductPostgres() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(configs.DB.PostgresDSN.Product)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	// Set connection pool configuration
	config.MaxConns = configs.DB.MaxConns
	config.MinConns = configs.DB.MinConns
	config.MaxConnIdleTime = configs.DB.MaxConnIdleTime

	pool, err := newPostgresWithConfig(config)
	if err != nil {
		return nil, err
	}

	log.Println("Product PostgreSQL connected")
	return pool, nil
}

// NewScorePostgres creates a new PostgreSQL connection pool for the score database
func NewScorePostgres() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(configs.DB.PostgresDSN.Score)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	// Set connection pool configuration
	config.MaxConns = configs.DB.MaxConns
	config.MinConns = configs.DB.MinConns
	config.MaxConnIdleTime = configs.DB.MaxConnIdleTime

	pool, err := newPostgresWithConfig(config)
	if err != nil {
		return nil, err
	}

	log.Println("Score PostgreSQL connected")
	return pool, nil
}

// newPostgresWithConfig creates a new PostgreSQL connection pool with the provided configuration
func newPostgresWithConfig(config *pgxpool.Config) (*pgxpool.Pool, error) {
	// Set connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), CONNECTION_TIMEOUT)
	defer cancel()

	pool, err := connectPostgresWithRetry(ctx, config)
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	return pool, nil
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
