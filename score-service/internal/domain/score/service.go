package score

import (
	"encoding/json"
	"log"
)

type ScoreService struct {
	repo ScoreRepository
}

// NewScoreService creates a new ScoreService instance
func NewScoreService(repo ScoreRepository) *ScoreService {
	return &ScoreService{
		repo: repo,
	}
}

// HandleBattleResults processes the battle results message
func (s *ScoreService) HandleBattleResults(msg []byte) error {
	var results BattleResults
	err := json.Unmarshal(msg, &results)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return err
	}

	var scores []Score
	for _, roundResult := range results.BattleResults {
		scores = append(scores, *NewScore(results.Game, roundResult.Round, roundResult.WinnerID, roundResult.LoserID))
	}

	err = s.repo.SaveMany(scores)
	if err != nil {
		log.Printf("Error saving scores: %v", err)
		return err
	}

	return nil
}
