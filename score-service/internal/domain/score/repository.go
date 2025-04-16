package score

type ScoreRepository interface {
	// SaveMany saves multiple scores to the database
	SaveMany(scores []Score) error
}
