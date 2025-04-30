package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oscarxxi/product-battle/bi-service/internal/domain/product"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Get() ([]product.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, `SELECT id, name FROM public.products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []product.Product{}
	for rows.Next() {
		var product product.Product
		if err := rows.Scan(&product.ID, &product.Name); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
