package app

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oscarxxi/product-battle/bi-service/internal/domain/product"
	"github.com/oscarxxi/product-battle/bi-service/internal/domain/score"
	"github.com/oscarxxi/product-battle/bi-service/internal/domain/widget"
	"github.com/oscarxxi/product-battle/bi-service/internal/infra/db"
	repository "github.com/oscarxxi/product-battle/bi-service/internal/infra/repositories/postgres"
	"github.com/oscarxxi/product-battle/bi-service/internal/interfaces/http"
)

type AppContainer struct {
	WidgetHandler *http.WidgetHandler
	productDB     *pgxpool.Pool
	scoreDB       *pgxpool.Pool
}

// NewAppContainer initializes the application container with dependencies.
func NewAppContainer() *AppContainer {
	productDB, err := db.NewProductPostgres()
	if err != nil {
		log.Fatalf("failed to connect to product database: %v", err)
	}

	scoreDB, err := db.NewScorePostgres()
	if err != nil {
		log.Fatalf("failed to connect to score database: %v", err)
	}

	productRepo := repository.NewProductRepository(productDB)
	scoreRepo := repository.NewScoreRepository(scoreDB)

	productService := product.NewProductService(productRepo)
	scoreService := score.NewScoreService(scoreRepo)

	widgetService := widget.NewWidgetService(productService, scoreService)

	widgetHandler := http.NewWidgetHandler(widgetService)

	return &AppContainer{
		WidgetHandler: widgetHandler,
		productDB:     productDB,
		scoreDB:       scoreDB,
	}
}

// Close cleans up the resources used by the application container.
func (c *AppContainer) Close() {
	c.productDB.Close()
	log.Println("Product database connection closed")

	c.scoreDB.Close()
	log.Println("Score database connection closed")

	log.Println("Application container closed")
}
