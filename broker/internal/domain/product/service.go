package product

import (
	product_proto "github.com/oscarxxi/product-battle/proto/product"
	"google.golang.org/grpc"
)

type ProductService struct {
	client product_proto.ProductServiceClient
	conn   *grpc.ClientConn
}
