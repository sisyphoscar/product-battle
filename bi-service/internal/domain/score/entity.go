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
