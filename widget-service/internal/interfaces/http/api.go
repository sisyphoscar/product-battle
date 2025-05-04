package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sisyphoscar/product-battle/bi-service/internal/domain/widget"
)

type WidgetHandler struct {
	widgetService *widget.WidgetService
}

// NewWidgetHandler initializes a new WidgetHandler
func NewWidgetHandler(widgetService *widget.WidgetService) *WidgetHandler {
	return &WidgetHandler{
		widgetService: widgetService,
	}
}

// Show handles the GET request for a specific widget.
func (h *WidgetHandler) Show(c *gin.Context) {
	widgetName := c.Param("widgetName")

	if widgetName != widget.PRODUCT_SCORE_WIDGET {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid widget name: " + widgetName})
		return
	}

	widget, err := h.widgetService.GetProductScoreWidget()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get widget"})
		return
	}

	c.JSON(http.StatusOK, widget)
}
