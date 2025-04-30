package product

type ProductService struct {
	repo ProductRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetIDNameMap retrieves a map of product IDs to product names.
func (s *ProductService) GetIDNameMap() (map[uint64]string, error) {
	products, err := s.get()
	if err != nil {
		return nil, err
	}

	productIdNameMap := make(map[uint64]string)
	for _, product := range products {
		productIdNameMap[product.ID] = product.Name
	}

	return productIdNameMap, nil
}

// get retrieves the list of products from the repository.
func (s *ProductService) get() ([]Product, error) {
	products, err := s.repo.Get()
	if err != nil {
		return nil, err
	}
	return products, nil
}
