package app

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sisyphoscar/product-battle/product-service/internal/domain/product"
	"github.com/sisyphoscar/product-battle/product-service/internal/infra/db"
	repository "github.com/sisyphoscar/product-battle/product-service/internal/infra/repositories/postgres"
)

type AppContainer struct {
	db             *pgxpool.Pool
	ProductService *product.ProductService
}

// NewAppContainer initializes the application container with dependencies.
func NewAppContainer() *AppContainer {
	db, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	productRepo := repository.NewProductRepository(db)
	productService := product.NewProductService(productRepo)

	return &AppContainer{
		db:             db,
		ProductService: productService,
	}
}

// Close cleans up the resources used by the application container.
func (c *AppContainer) Close() {
	c.db.Close()
	log.Println("Database connection closed")
	log.Println("Application container closed")
}
