package repository

import (
	"context"
	"product/internal/domain/product"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]product.Product, error) {
	// 5 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, "SELECT id, name, description, image_url, price, created_at, updated_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []product.Product
	for rows.Next() {
		var p product.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.ImageURL, &p.Price, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
