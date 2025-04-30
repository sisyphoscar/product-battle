package widget

import (
	"context"
	"encoding/json"
	"time"

	widget_proto "github.com/oscarxxi/product-battle/proto/widget"
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
func (w *WidgetService) GetWidget(name string) (*Widget, error) {
	req := &widget_proto.WidgetRequest{
		Name: name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := w.client.GetWidget(ctx, req)
	if err != nil {
		return nil, err
	}

	var stats interface{}
	if err := json.Unmarshal([]byte(res.Stats), &stats); err != nil {
		return nil, err
	}

	widget := Widget{
		Name:  res.Name,
		Stats: stats,
	}

	return &widget, nil
}
