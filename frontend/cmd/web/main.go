package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./resources/static")
	router.LoadHTMLGlob("resources/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"productEndpoint": "http://localhost:8081",
		})
	})

	router.Run(":8080")
}
