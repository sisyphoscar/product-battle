package main

import (
	"broker/internal/app"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app.LoadConfig()

	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Starting server on", app.App.URL)

	router.Run(app.App.URL)
}
