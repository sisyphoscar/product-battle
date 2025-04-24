package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/domain/product"
	"github.com/oscarxxi/product-battle/broker/internal/interfaces/http/dto"
)

type ProductHandler struct {
	productService *product.ProductService
}

// NewProductHandler initializes a new ProductHandler
func NewProductHandler(productService *product.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// GetAllProducts handles the GET request to fetch all products
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to fetch products",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    products,
	})
}
