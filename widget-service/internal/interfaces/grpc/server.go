package grpc

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/sisyphoscar/product-battle/bi-service/internal/app/configs"
	"github.com/sisyphoscar/product-battle/bi-service/internal/domain/widget"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	widget_proto "github.com/sisyphoscar/product-battle-proto/widget"
)

type WidgetServer struct {
	widget_proto.UnimplementedWidgetServiceServer
	service *widget.WidgetService
}

// Listen starts the gRPC server and listens for incoming requests
func Listen(service *widget.WidgetService) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+configs.App.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	widget_proto.RegisterWidgetServiceServer(s, &WidgetServer{service: service})

	log.Printf("Listening and serving gRPC on %s", configs.App.GRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

// GetWidget handles the gRPC request to get a widget by name
func (s *WidgetServer) GetWidget(ctx context.Context, req *widget_proto.WidgetRequest) (*widget_proto.WidgetResponse, error) {
	widgetName := req.GetName()
	if widgetName != widget.PRODUCT_SCORE_WIDGET {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid widget name: %s", widgetName)
	}

	widget, err := s.service.GetProductScoreWidget()
	if err != nil {
		return nil, err
	}

	statsJSON, err := json.Marshal(widget.Stats)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to marshal stats to JSON: %v", err)
	}

	return &widget_proto.WidgetResponse{
		Name:  widgetName,
		Stats: string(statsJSON),
	}, nil
}
