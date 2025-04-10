package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/product-service/internal/app"
	"github.com/oscarxxi/product-battle/product-service/internal/app/configs"
	"github.com/oscarxxi/product-battle/product-service/internal/interfaces/grpc"
)

func main() {
	configs.LoadConfig()

	container := app.NewAppContainer()
	defer container.Close()

	go grpc.Listen(container.ProductService)

	httpListen()
}

func httpListen() {
	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Starting HTTP server on", configs.App.URL)

	router.Run(configs.App.URL)
}
