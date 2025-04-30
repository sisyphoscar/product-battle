package widget

const (
	PRODUCT_SCORE_WIDGET = "product-score"
)

type Widget struct {
	Name  string      `json:"name"`
	Stats interface{} `json:"stats"`
}
