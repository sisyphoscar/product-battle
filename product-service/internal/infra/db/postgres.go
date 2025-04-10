package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oscarxxi/product-battle/product-service/internal/app/configs"
)

// NewPostgres initializes a new PostgreSQL connection pool.
func NewPostgres() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(configs.DB.PostgresDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	// Set connection pool configuration
	config.MaxConns = configs.DB.MaxConns
	config.MinConns = configs.DB.MinConns
	config.MaxConnIdleTime = configs.DB.MaxConnIdleTime

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

	log.Println("Database connected")
	return pool, nil
}
