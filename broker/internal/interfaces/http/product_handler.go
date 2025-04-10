package http_interface

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/domain/product"
)

type ProductHandler struct {
	productService *product.ProductService
}

type GetAllProductsResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []product.Product `json:"data"`
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
