package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/app/configs"
)

func main() {
	configs.LoadConfig()

	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Starting server on", configs.App.URL)

	router.Run(configs.App.URL)
}
