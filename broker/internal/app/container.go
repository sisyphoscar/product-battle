package app

import (
	"log"

	"github.com/sisyphoscar/product-battle/broker/internal/app/configs"
	"github.com/sisyphoscar/product-battle/broker/internal/domain/product"
	"github.com/sisyphoscar/product-battle/broker/internal/domain/widget"
	"github.com/sisyphoscar/product-battle/broker/internal/infra/messaging"
	handlers "github.com/sisyphoscar/product-battle/broker/internal/interfaces/http/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AppContainer struct {
	// Handlers
	ProductHandler *handlers.ProductHandler
	BattleHandler  *handlers.BattleHandler
	WidgetHandler  *handlers.WidgetHandler
	// Services
	productService *product.ProductService
	widgetService  *widget.WidgetService
	// Infrastructure
	productGRPCConn *grpc.ClientConn
	widgetGRPCConn  *grpc.ClientConn
	rabbitMQ        *messaging.RabbitMQ
}

// NewAppContainer initializes the application container with dependencies.
func NewAppContainer() *AppContainer {
	var ac AppContainer

	ac.loadInfra()
	ac.loadServices()
	ac.loadHandlers()

	return &ac
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

// loadInfra initializes the infrastructure components.
func (ac *AppContainer) loadInfra() {
	rabbitMQ, err := messaging.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	productGRPCConn, err := grpc.NewClient(
		configs.Endpoint.ProductService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	widgetGRPCConn, err := grpc.NewClient(
		configs.Endpoint.WidgetService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Widget gRPC server: %v", err)
	}

	ac.rabbitMQ = rabbitMQ
	ac.productGRPCConn = productGRPCConn
	ac.widgetGRPCConn = widgetGRPCConn
}

// loadServices initializes the services
func (ac *AppContainer) loadServices() {
	ps := product.NewProductService(ac.productGRPCConn)
	ws := widget.NewWidgetService(ac.widgetGRPCConn)

	ac.productService = ps
	ac.widgetService = ws
}

// loadHandlers initializes the HTTP handlers
func (ac *AppContainer) loadHandlers() {
	ph := handlers.NewProductHandler(ac.productService)
	bh := handlers.NewBattleHandler(ac.rabbitMQ)
	wh := handlers.NewWidgetHandler(ac.widgetService)

	ac.ProductHandler = ph
	ac.BattleHandler = bh
	ac.WidgetHandler = wh
}
