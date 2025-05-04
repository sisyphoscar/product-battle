package widget

import (
	"github.com/sisyphoscar/product-battle/bi-service/internal/domain/product"
	"github.com/sisyphoscar/product-battle/bi-service/internal/domain/score"
)

type WidgetService struct {
	productService *product.ProductService
	scoreService   *score.ScoreService
}

// NewWidgetService initializes a new WidgetService
func NewWidgetService(productService *product.ProductService, scoreService *score.ScoreService) *WidgetService {
	return &WidgetService{
		productService: productService,
		scoreService:   scoreService,
	}
}

// GetProductScoreWidget retrieves the product score widget data.
func (s *WidgetService) GetProductScoreWidget() (Widget, error) {
	stats, err := s.scoreService.CountScoreGroupByWinner()
	if err != nil {
		return Widget{}, err
	}

	productIdNameMap, err := s.productService.GetIDNameMap()
	if err != nil {
		return Widget{}, err
	}

	// update the product names in the stats
	for i, stat := range stats {
		productName, ok := productIdNameMap[stat.ProductID]
		if ok {
			stats[i].ProductName = productName
		} else {
			stats[i].ProductName = "Unknown"
		}
	}

	return Widget{
		Name:  PRODUCT_SCORE_WIDGET,
		Stats: stats,
	}, nil
}
