package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sisyphoscar/product-battle/broker/internal/app/configs"
	"github.com/sisyphoscar/product-battle/broker/internal/infra/messaging"
	"github.com/sisyphoscar/product-battle/broker/internal/interfaces/http/dto"
)

type BattleHandler struct {
	rabbitMQ *messaging.RabbitMQ
}

// NewBattleHandler initializes a new BattleHandler
func NewBattleHandler(rabbitMQ *messaging.RabbitMQ) *BattleHandler {
	return &BattleHandler{
		rabbitMQ: rabbitMQ,
	}
}

type SubmitProductBattleRequest struct {
	Game          string         `json:"game" binding:"required"`
	BattleResults []BattleResult `json:"roundResults" binding:"required,dive"`
}

type BattleResult struct {
	Round    int    `json:"round" binding:"required,gt=0"`
	WinnerID uint64 `json:"winnerId" binding:"required"`
	LoserID  uint64 `json:"loserId" binding:"required"`
}

// SubmitProductBattle handles the POST request to submit product battle results
func (h *BattleHandler) SubmitProductBattle(c *gin.Context) {
	var results SubmitProductBattleRequest
	if err := c.ShouldBindJSON(&results); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
		return
	}

	resultsJSON, err := json.Marshal(results)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to marshal JSON",
			Error:   err.Error(),
		})
		return
	}

	// publish the results to RabbitMQ
	err = h.rabbitMQ.Publish(configs.Queue.BattleScoreQueue, string(resultsJSON))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to publish message to RabbitMQ",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, dto.SuccessResponse{
		Status:  http.StatusAccepted,
		Message: "success",
		Data:    results,
	})
}
