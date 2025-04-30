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

type ScoreGroupByWinnerResult struct {
	ProductID   uint64 `json:"-"`
	ProductName string `json:"productName"`
	Score       int    `json:"score"`
}
