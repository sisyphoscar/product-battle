package score

type ScoreService struct {
	repo ScoreRepository
}

// NewScoreService creates a new ScoreService instance
func NewScoreService(repo ScoreRepository) *ScoreService {
	return &ScoreService{
		repo: repo,
	}
}
