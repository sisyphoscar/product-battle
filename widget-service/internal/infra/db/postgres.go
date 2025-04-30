package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oscarxxi/product-battle/bi-service/internal/app/configs"
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
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := newPostgresWithConfig(config)
	if err != nil {
		return nil, err
	}

	log.Println("Score PostgreSQL connected")
	return pool, nil
}

func newPostgresWithConfig(config *pgxpool.Config) (*pgxpool.Pool, error) {
	// Set connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a new connection pool
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
