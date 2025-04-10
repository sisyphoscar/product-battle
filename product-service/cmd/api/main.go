package main

import (
	"log"
	"net/http"
	"product/internal/app"
	"product/internal/domain/product"
	"product/internal/infra/db"
	repository "product/internal/infra/repositories/postgres"
	interface_http "product/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app.LoadConfig()

	db, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	productService := product.NewProductService(productRepo)
	productHandler := interface_http.NewProductHandler(productService)

	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/products", productHandler.GetProducts)
	}

	log.Println("Starting server on", app.App.URL)

	router.Run(app.App.URL)
}
