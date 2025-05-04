package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sisyphoscar/product-battle/broker/internal/domain/widget"
	"github.com/sisyphoscar/product-battle/broker/internal/interfaces/http/dto"
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
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid widget name",
			Error:   "Invalid widget name: " + widgetName,
		})
		return
	}

	widgetData, err := h.WidgetService.GetWidget(widgetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get widget data",
			Error:   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, widgetData)
}
