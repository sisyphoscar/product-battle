package interface_http

import (
	"net/http"
	"product/internal/domain/product"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service product.ProductService
}

// NewProductHandler creates a new instance of ProductHandler
func NewProductHandler(service *product.ProductService) *ProductHandler {
	return &ProductHandler{service: *service}
}

// GetProducts handles the GET request to fetch all products
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
