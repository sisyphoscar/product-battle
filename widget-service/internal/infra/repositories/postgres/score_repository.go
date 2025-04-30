package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oscarxxi/product-battle/bi-service/internal/domain/score"
)

type ScoreRepository struct {
	db *pgxpool.Pool
}

// NewScoreRepository creates a new instance of ScoreRepository
func NewScoreRepository(db *pgxpool.Pool) *ScoreRepository {
	return &ScoreRepository{db: db}
}

func (r *ScoreRepository) CountScoreGroupByWinner() ([]score.ScoreGroupByWinnerResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, `
		SELECT
			winner_id AS product_id,
			COUNT(*) AS score
		FROM
			score_logs
		GROUP BY
			winner_id
		ORDER BY
			score DESC;
	`)
	if err != nil {
		return []score.ScoreGroupByWinnerResult{}, err
	}
	defer rows.Close()

	results := []score.ScoreGroupByWinnerResult{}

	for rows.Next() {
		var r score.ScoreGroupByWinnerResult
		if err := rows.Scan(&r.ProductID, &r.Score); err != nil {
			return []score.ScoreGroupByWinnerResult{}, err
		}
		results = append(results, r)
	}
	if err := rows.Err(); err != nil {
		return []score.ScoreGroupByWinnerResult{}, err
	}

	return results, nil
}
