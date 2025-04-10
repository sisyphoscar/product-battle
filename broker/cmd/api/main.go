package main

import (
	"broker/internal/app/configs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
