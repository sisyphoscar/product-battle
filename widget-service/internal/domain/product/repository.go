package product

type ProductRepository interface {
	// Get retrieves the list of products.
	Get() ([]Product, error)
}
