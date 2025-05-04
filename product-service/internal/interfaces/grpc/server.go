package grpc

import (
	"context"
	"log"
	"net"

	"github.com/sisyphoscar/product-battle/product-service/internal/app/configs"
	"github.com/sisyphoscar/product-battle/product-service/internal/domain/product"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	product_proto "github.com/sisyphoscar/product-battle-proto/product"
)

type ProductServer struct {
	product_proto.UnimplementedProductServiceServer
	service *product.ProductService
}

// Listen starts the gRPC server and listens for incoming requests
func Listen(service *product.ProductService) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+configs.App.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	product_proto.RegisterProductServiceServer(s, &ProductServer{service: service})

	log.Printf("Listening and serving gRPC on %s", configs.App.GRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

// GetAllProducts handles the gRPC request to fetch all products
func (s *ProductServer) GetAllProducts(context.Context, *emptypb.Empty) (*product_proto.ProductList, error) {
	products, err := s.service.GetAll()
	if err != nil {
		return nil, err
	}

	productList := &product_proto.ProductList{}
	for _, p := range products {
		productList.Products = append(productList.Products, &product_proto.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			ImageUrl:    p.ImageURL,
		})
	}

	return productList, nil
}
