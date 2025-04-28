package score

type ScoreRepository interface {
	// CountScoreGroupByProduct counts the scores grouped by winner of the round.
	CountScoreGroupByWinner() ([]ScoreGroupByWinnerResult, error)
}
