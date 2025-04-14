package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/app/configs"
	"github.com/oscarxxi/product-battle/broker/internal/infra/messaging"
	"github.com/oscarxxi/product-battle/broker/internal/interfaces/http/helpers"
)

type BattleHandler struct {
	RabbitMQ *messaging.RabbitMQ
}

// NewBattleHandler initializes a new BattleHandler
func NewBattleHandler(rabbitMQ *messaging.RabbitMQ) *BattleHandler {
	return &BattleHandler{
		RabbitMQ: rabbitMQ,
	}
}

type SubmitProductBattleRequest struct {
	SeasonID      string         `json:"seasonId" binding:"required"`
	BattleResults []BattleResult `json:"roundResults" binding:"required,dive"`
}

type BattleResult struct {
	Round    int    `json:"round" binding:"required,gt=0"`
	WinnerID string `json:"winnerId" binding:"required"`
	LoserID  string `json:"loserId" binding:"required"`
}

// SubmitProductBattle handles the POST request to submit product battle results
func (h *BattleHandler) SubmitProductBattle(c *gin.Context) {
	var results SubmitProductBattleRequest
	if err := c.ShouldBindJSON(&results); err != nil {
		helpers.ErrorResponse(c, err, helpers.ResponseOptions{
			Status:  http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	resultsJSON, err := json.Marshal(results)
	if err != nil {
		helpers.ErrorResponse(c, err, helpers.ResponseOptions{
			Status:  http.StatusInternalServerError,
			Message: "Failed to marshal JSON",
		})
		return
	}

	// publish the results to RabbitMQ
	err = h.RabbitMQ.Publish(configs.Queue.BattleResultQueue, string(resultsJSON))
	if err != nil {
		helpers.ErrorResponse(c, err, helpers.ResponseOptions{
			Status:  http.StatusInternalServerError,
			Message: "Failed to publish message",
		})
		return
	}

	helpers.SuccessResponse(c, results)
}
