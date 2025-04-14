package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/app"
	"github.com/oscarxxi/product-battle/broker/internal/app/configs"
	"github.com/oscarxxi/product-battle/broker/internal/infra/messaging"
	http_interface "github.com/oscarxxi/product-battle/broker/internal/interfaces/http"
)

func main() {
	configs.LoadConfig()

	appContainer := app.NewAppContainer()
	defer appContainer.Close()

	rabbitMQ, err := messaging.NewRabbitMQ("amqp://admin:admin@localhost:5672")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// HTTP server setup
	router := gin.Default()
	router = http_interface.SetApiRoutes(router, appContainer.ProductHandler)

	server := &http.Server{
		Addr:    configs.App.URL,
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
}
