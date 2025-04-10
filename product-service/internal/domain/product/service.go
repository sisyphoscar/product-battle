package product

type ProductService struct {
	repo ProductRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Get all products
func (s *ProductService) GetAll() ([]Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
