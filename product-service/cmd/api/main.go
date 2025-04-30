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
	"github.com/oscarxxi/product-battle/product-service/internal/app"
	"github.com/oscarxxi/product-battle/product-service/internal/app/configs"
	"github.com/oscarxxi/product-battle/product-service/internal/interfaces/grpc"
)

func main() {
	configs.LoadConfig()

	appContainer := app.NewAppContainer()
	defer appContainer.Close()

	// gRPC server setup
	go grpc.Listen(appContainer.ProductService)

	// HTTP server setup
	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
}
