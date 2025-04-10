package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/app/configs"
	"github.com/oscarxxi/product-battle/broker/internal/domain/product"
)

func main() {
	configs.LoadConfig()

	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/products", func(c *gin.Context) {
		service := product.NewProductService()
		defer service.Close()

		products, err := service.GetAllProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	log.Println("Starting server on", configs.App.URL)

	router.Run(configs.App.URL)
}
