package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/bi-service/internal/domain/product"
	"github.com/oscarxxi/product-battle/bi-service/internal/domain/score"
)

type BIHandler struct {
	productService *product.ProductService
	scoreService   *score.ScoreService
}

// NewBIHandler initializes a new BIHandler
func NewBIHandler(productService *product.ProductService, scoreService *score.ScoreService) *BIHandler {
	return &BIHandler{
		productService: productService,
		scoreService:   scoreService,
	}
}

func (h *BIHandler) GetProductScoreStats(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
