package product

import (
	"context"
	"time"

	product_proto "github.com/sisyphoscar/product-battle-proto/product"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct {
	client product_proto.ProductServiceClient
	conn   *grpc.ClientConn
}

// NewProductService initializes a new ProductService
func NewProductService(conn *grpc.ClientConn) *ProductService {
	return &ProductService{
		client: product_proto.NewProductServiceClient(conn),
		conn:   conn,
	}
}

// GetAllProducts retrieves all products from the product service
func (s *ProductService) GetAllProducts() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := s.client.GetAllProducts(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	var products []Product
	for _, p := range resp.Products {
		products = append(products, Product{
			ID:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			ImageURL:    p.ImageUrl,
		})
	}

	return products, nil
}
