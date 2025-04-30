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

// CountScoreGroupByWinner counts the scores grouped by winner of the round.
func (s *ScoreService) CountScoreGroupByWinner() ([]ScoreGroupByWinnerResult, error) {
	scores, err := s.repo.CountScoreGroupByWinner()
	if err != nil {
		return []ScoreGroupByWinnerResult{}, err
	}
	return scores, nil
}
