package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oscarxxi/product-battle/web/internal/app/configs"
)

func main() {
	configs.LoadConfig()

	router := gin.Default()

	router.Static("/static", "./resources/static")
	router.LoadHTMLGlob("resources/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"game":           uuid.New().String(),
			"brokerEndpoint": configs.Endpoint.Broker,
		})
	})

	router.Run(configs.App.URL)
}
