package score

import "time"

type Score struct {
	ID        string
	Game      string
	Round     int
	WinnerID  uint64
	LoserID   uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BattleResults struct {
	Game          string        `json:"game" binding:"required"`
	BattleResults []RoundResult `json:"roundResults" binding:"required,dive"`
}

type RoundResult struct {
	Round    int    `json:"round" binding:"required,gt=0"`
	WinnerID uint64 `json:"winnerId" binding:"required"`
	LoserID  uint64 `json:"loserId" binding:"required"`
}

// NewScore creates a new Score instance
func NewScore(game string, round int, winnerID uint64, loserID uint64) *Score {
	now := time.Now()

	return &Score{
		Game:      game,
		Round:     round,
		WinnerID:  winnerID,
		LoserID:   loserID,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
