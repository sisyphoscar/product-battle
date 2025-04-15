package score

import (
	"encoding/json"
	"log"
)

type ScoreService struct {
}

func NewScoreService() *ScoreService {
	return &ScoreService{}
}

type ProductBattleResults struct {
	SeasonID      string        `json:"seasonId" binding:"required"`
	BattleResults []RoundResult `json:"roundResults" binding:"required,dive"`
}

type RoundResult struct {
	Round    int    `json:"round" binding:"required,gt=0"`
	WinnerID string `json:"winnerId" binding:"required"`
	LoserID  string `json:"loserId" binding:"required"`
}

// HandleBattleResults processes the battle results message
func (s *ScoreService) HandleBattleResults(msg []byte) error {
	var results ProductBattleResults
	err := json.Unmarshal(msg, &results)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return err
	}

	// Process the message here...
	log.Println("Processing battle results:", results)

	return nil
}
