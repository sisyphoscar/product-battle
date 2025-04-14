package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/app"
)

// Api sets up the API routes for the application
func SetApiRoutes(router *gin.Engine, ac *app.AppContainer) *gin.Engine {
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/products", ac.ProductHandler.GetAllProducts)
		apiGroup.POST("/product-battle/submit", ac.BattleHandler.SubmitProductBattle)
	}

	return router
}
