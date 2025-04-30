package app

import (
	"log"

	"github.com/oscarxxi/product-battle/broker/internal/app/configs"
	"github.com/oscarxxi/product-battle/broker/internal/domain/product"
	"github.com/oscarxxi/product-battle/broker/internal/domain/widget"
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
	WidgetService   *widget.WidgetService
	widgetGRPCConn  *grpc.ClientConn
	WidgetHandler   *handlers.WidgetHandler
}

// NewAppContainer initializes the application container with dependencies.
func NewAppContainer() *AppContainer {
	// load infra
	// load repository
	// load service
	// load handler
	rabbitMQ, err := messaging.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	battleHandler := handlers.NewBattleHandler(rabbitMQ)

	productGRPCConn, err := grpc.NewClient(
		configs.Endpoint.ProductService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	productService := product.NewProductService(productGRPCConn)
	productHandler := handlers.NewProductHandler(productService)

	widgetGRPCConn, err := grpc.NewClient(
		configs.Endpoint.WidgetService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Widget gRPC server: %v", err)
	}

	widgetService := widget.NewWidgetService(widgetGRPCConn)
	widgetHandler := handlers.NewWidgetHandler(widgetService)

	return &AppContainer{
		ProductHandler:  productHandler,
		BattleHandler:   battleHandler,
		rabbitMQ:        rabbitMQ,
		productGRPCConn: productGRPCConn,
		WidgetService:   widgetService,
		widgetGRPCConn:  widgetGRPCConn,
		WidgetHandler:   widgetHandler,
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
		log.Printf("Failed to close product gRPC connection: %v", err)
	} else {
		log.Println("Product gRPC connection closed")
	}

	err = c.widgetGRPCConn.Close()
	if err != nil {
		log.Printf("Failed to close Widget gRPC connection: %v", err)
	} else {
		log.Println("Widget gRPC connection closed")
	}

	log.Println("Application container closed")
}
