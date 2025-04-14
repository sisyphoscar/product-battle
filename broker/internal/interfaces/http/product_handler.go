package http_interface

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/app/configs"
	"github.com/oscarxxi/product-battle/broker/internal/domain/product"
	"github.com/oscarxxi/product-battle/broker/internal/infra/messaging"
)

type ProductHandler struct {
	ProductService *product.ProductService
	RabbitMQ       *messaging.RabbitMQ
}

type GetAllProductsResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []product.Product `json:"data"`
}

// NewProductHandler initializes a new ProductHandler
func NewProductHandler(productService *product.ProductService, rabbitMQ *messaging.RabbitMQ) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
		RabbitMQ:       rabbitMQ,
	}
}

// GetAllProducts handles the GET request to fetch all products
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.ProductService.GetAllProducts()
	if err != nil {
		res := GetAllProductsResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to fetch products",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := GetAllProductsResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    products,
	}
	c.JSON(http.StatusOK, res)
}

// SettleProductScore handles the POST request to settle product scores
func (h *ProductHandler) SettleProductScore(c *gin.Context) {
	type RoundResult struct {
		Round    int    `json:"round"`
		WinnerID string `json:"winner_id"`
		LoserID  string `json:"loser_id"`
	}

	type ScoreResults struct {
		SeasonID      string        `json:"seasonId"`
		BattleResults []RoundResult `json:"roundResults"`
	}

	scoreResults := ScoreResults{
		SeasonID: "abc123",
		BattleResults: []RoundResult{
			{Round: 1, WinnerID: "b1", LoserID: "b2"},
			{Round: 2, WinnerID: "b1", LoserID: "b3"},
		},
	}

	scoreResultsJSON, err := json.Marshal(scoreResults)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	// Publish the message to RabbitMQ
	err = h.RabbitMQ.Publish(
		configs.Queue.ScoreSettlementQueue,
		string(scoreResultsJSON),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    scoreResults,
	})
}
