package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ScoreRepository struct {
	db *pgxpool.Pool
}

// NewScoreRepository creates a new instance of ScoreRepository
func NewScoreRepository(db *pgxpool.Pool) *ScoreRepository {
	return &ScoreRepository{db: db}
}
