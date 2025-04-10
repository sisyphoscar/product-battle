package main

import (
	"log"
	"net/http"
	"product/internal/app/configs"
	"product/internal/domain/product"
	"product/internal/infra/db"
	repository "product/internal/infra/repositories/postgres"
	"product/internal/interfaces/grpc"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	db, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	productService := product.NewProductService(productRepo)

	go grpc.Listen(productService)

	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Starting HTTP server on", configs.App.URL)

	router.Run(configs.App.URL)
}
