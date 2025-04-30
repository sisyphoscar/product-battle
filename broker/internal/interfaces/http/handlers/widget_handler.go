package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/domain/widget"
)

type WidgetHandler struct {
	WidgetService *widget.WidgetService
}

// NewWidgetHandler initializes a new WidgetHandler
func NewWidgetHandler(widgetService *widget.WidgetService) *WidgetHandler {
	return &WidgetHandler{
		WidgetService: widgetService,
	}
}

// Show handles the request to show a widget by its name
func (h *WidgetHandler) Show(c *gin.Context) {
	widgetName := c.Param("widgetName")

	if widgetName != widget.PRODUCT_SCORE_WIDGET {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid widget name: " + widgetName})
		return
	}

	widgetData, err := h.WidgetService.GetWidget(widgetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get widget data"})
		return
	}

	c.JSON(http.StatusOK, widgetData)
}
