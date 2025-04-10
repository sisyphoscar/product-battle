package product

type ProductRepository interface {
	// Get all products
	GetAll() ([]Product, error)
}
