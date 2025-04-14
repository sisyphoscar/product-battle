package http_interface

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Api sets up the API routes for the application
func SetApiRoutes(router *gin.Engine, handler *ProductHandler) *gin.Engine {
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/products", handler.GetAllProducts)
	}

	return router
}
