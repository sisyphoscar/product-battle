package widget

type Widget struct {
	Name  string      `json:"name"`
	Stats interface{} `json:"stats"`
}

const (
	PRODUCT_SCORE_WIDGET = "product-score"
)
