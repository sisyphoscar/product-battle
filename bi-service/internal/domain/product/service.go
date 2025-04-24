package product

type ProductService struct {
	repo ProductRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
