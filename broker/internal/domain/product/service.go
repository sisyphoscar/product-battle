package product

import (
	"context"
	"log"

	product_proto "github.com/oscarxxi/product-battle/proto/product"
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

// Close closes the gRPC connection
func (s *ProductService) Close() {
	err := s.conn.Close()
	if err != nil {
		log.Printf("Failed to close gRPC connection: %v", err)
	}
	log.Println("ProductService connection closed")
}

// GetAllProducts retrieves all products from the product service
func (c *ProductService) GetAllProducts() ([]Product, error) {
	resp, err := c.client.GetAllProducts(context.Background(), &emptypb.Empty{})
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
