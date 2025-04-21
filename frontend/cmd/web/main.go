package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("resources/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/game", func(c *gin.Context) {
		c.HTML(http.StatusOK, "game.html", gin.H{
			"productEndpoint": "http://localhost:8081",
		})
	})

	router.Run(":8080")
}
