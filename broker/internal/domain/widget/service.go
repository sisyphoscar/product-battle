package widget

import (
	"context"
	"encoding/json"
	"time"

	widget_proto "github.com/sisyphoscar/product-battle-proto/widget"
	"google.golang.org/grpc"
)

type WidgetService struct {
	client widget_proto.WidgetServiceClient
	conn   *grpc.ClientConn
}

// NewWidgetService creates a new WidgetService instance with the provided gRPC connection.
func NewWidgetService(conn *grpc.ClientConn) *WidgetService {
	return &WidgetService{
		client: widget_proto.NewWidgetServiceClient(conn),
		conn:   conn,
	}
}

// GetWidget retrieves the widget data by name.
func (s *WidgetService) GetWidget(name string) (*Widget, error) {
	req := &widget_proto.WidgetRequest{
		Name: name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := s.client.GetWidget(ctx, req)
	if err != nil {
		return nil, err
	}

	var stats interface{}
	if err := json.Unmarshal([]byte(resp.Stats), &stats); err != nil {
		return nil, err
	}

	widget := Widget{
		Name:  resp.Name,
		Stats: stats,
	}

	return &widget, nil
}
