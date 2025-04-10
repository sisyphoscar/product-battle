package product

type ProductRepository interface {
	// Get all products
	Get() ([]Product, error)
}
