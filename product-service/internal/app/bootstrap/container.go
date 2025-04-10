package bootstrap

import (
	"log"
	"product/internal/domain/product"
	"product/internal/infra/db"
	repository "product/internal/infra/repositories/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AppContainer struct {
	DB             *pgxpool.Pool
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
		DB:             db,
		ProductService: productService,
	}
}

// Close cleans up the resources used by the application container.
func (c *AppContainer) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
