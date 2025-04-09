package main

import (
	"log"
	"net/http"
	"product/internal/app"
	"product/internal/infra/db"

	"github.com/gin-gonic/gin"
)

func main() {
	app.LoadConfig()

	db, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Starting server on", app.App.URL)

	router.Run(app.App.URL)
}
