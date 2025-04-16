package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oscarxxi/product-battle/score-service/internal/domain/score"
)

type ScoreRepository struct {
	db *pgxpool.Pool
}

// NewScoreRepository creates a new instance of ScoreRepository
func NewScoreRepository(db *pgxpool.Pool) *ScoreRepository {
	return &ScoreRepository{db: db}
}

func (r *ScoreRepository) SaveMany(scores []score.Score) error {
	if len(scores) == 0 {
		return nil
	}

	// prepare arguments for query binding
	values := []string{}
	args := []interface{}{}
	for row, sc := range scores {
		values = append(values, "($"+strconv.Itoa(row*4+1)+", $"+strconv.Itoa(row*4+2)+", $"+strconv.Itoa(row*4+3)+", $"+strconv.Itoa(row*4+4)+")")
		args = append(args, sc.Game, sc.Round, sc.WinnerID, sc.LoserID)
	}

	query := fmt.Sprintf(
		"INSERT INTO score_logs (game, round, winner_id, loser_id) VALUES %s",
		strings.Join(values, ","),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
