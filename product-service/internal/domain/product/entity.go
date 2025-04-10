package product

import "time"

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
