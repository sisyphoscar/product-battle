package app

import (
	"log"

	"github.com/oscarxxi/product-battle/broker/internal/app/configs"
	"github.com/oscarxxi/product-battle/broker/internal/domain/product"
	"github.com/oscarxxi/product-battle/broker/internal/infra/messaging"
	handlers "github.com/oscarxxi/product-battle/broker/internal/interfaces/http/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AppContainer struct {
	ProductHandler  *handlers.ProductHandler
	BattleHandler   *handlers.BattleHandler
	rabbitMQ        *messaging.RabbitMQ
	productGRPCConn *grpc.ClientConn
}

// NewAppContainer initializes the application container with dependencies.
func NewAppContainer() *AppContainer {
	rabbitMQ, err := messaging.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	battleHandler := handlers.NewBattleHandler(rabbitMQ)

	conn, err := grpc.NewClient(
		configs.Endpoint.ProductService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	productService := product.NewProductService(conn)
	productHandler := handlers.NewProductHandler(productService)

	return &AppContainer{
		ProductHandler:  productHandler,
		BattleHandler:   battleHandler,
		rabbitMQ:        rabbitMQ,
		productGRPCConn: conn,
	}
}

// Close cleans up the resources used by the application container.
func (c *AppContainer) Close() {
	err := c.rabbitMQ.Close()
	if err != nil {
		log.Printf("Failed to close RabbitMQ connection: %v", err)
	} else {
		log.Println("RabbitMQ connection closed")
	}

	err = c.productGRPCConn.Close()
	if err != nil {
		log.Printf("Failed to close gRPC connection: %v", err)
	} else {
		log.Println("gRPC connection closed")
	}

	log.Println("Application container closed")
}
