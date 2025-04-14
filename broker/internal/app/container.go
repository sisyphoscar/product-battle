package app

import (
	"log"

	"github.com/oscarxxi/product-battle/broker/internal/domain/product"
	"github.com/oscarxxi/product-battle/broker/internal/infra/messaging"
	handlers "github.com/oscarxxi/product-battle/broker/internal/interfaces/http/handlers"
)

type AppContainer struct {
	ProductHandler *handlers.ProductHandler
	BattleHandler  *handlers.BattleHandler
}

// NewAppContainer initializes the application container with dependencies.
func NewAppContainer() *AppContainer {
	rabbitMQ, err := messaging.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	battleHandler := handlers.NewBattleHandler(rabbitMQ)

	productService := product.NewProductService()
	productHandler := handlers.NewProductHandler(productService)

	return &AppContainer{
		ProductHandler: productHandler,
		BattleHandler:  battleHandler,
	}
}

// Close cleans up the resources used by the application container.
func (c *AppContainer) Close() {
	c.ProductHandler.ProductService.Close()
	c.BattleHandler.RabbitMQ.Close()

	log.Println("Application container closed")
}
