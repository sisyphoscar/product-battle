package app

import (
	"log"

	"github.com/oscarxxi/product-battle/broker/internal/domain/product"
	http_interface "github.com/oscarxxi/product-battle/broker/internal/interfaces/http"
)

type AppContainer struct {
	ProductService *product.ProductService
	ProductHandler *http_interface.ProductHandler
}

// NewAppContainer initializes the application container with dependencies.
func NewAppContainer() *AppContainer {
	productService := product.NewProductService()
	productHandler := http_interface.NewProductHandler(productService)

	return &AppContainer{
		ProductService: productService,
		ProductHandler: productHandler,
	}
}

// Close cleans up the resources used by the application container.
func (c *AppContainer) Close() {
	c.ProductService.Close()
	log.Println("Application container closed")
}
