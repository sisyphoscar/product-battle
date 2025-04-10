package product

import "time"

type Product struct {
	ID          uint64
	Name        string
	Description string
	ImageURL    string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
