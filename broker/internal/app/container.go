package app

import (
	"log"

	"github.com/oscarxxi/product-battle/broker/internal/domain/product"
	"github.com/oscarxxi/product-battle/broker/internal/infra/messaging"
	http_interface "github.com/oscarxxi/product-battle/broker/internal/interfaces/http"
)

type AppContainer struct {
	ProductHandler *http_interface.ProductHandler
}

// NewAppContainer initializes the application container with dependencies.
func NewAppContainer() *AppContainer {
	rabbitMQ, err := messaging.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	productService := product.NewProductService()
	productHandler := http_interface.NewProductHandler(productService, rabbitMQ)

	return &AppContainer{
		ProductHandler: productHandler,
	}
}

// Close cleans up the resources used by the application container.
func (c *AppContainer) Close() {
	c.ProductHandler.ProductService.Close()
	c.ProductHandler.RabbitMQ.Close()
	log.Println("Application container closed")
}
