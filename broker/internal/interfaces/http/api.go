package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sisyphoscar/product-battle/broker/internal/app"
	middleware "github.com/sisyphoscar/product-battle/broker/internal/interfaces/http/middlewares"
)

// SetApiRoutes sets up the API routes for the application.
func SetApiRoutes(r *gin.Engine, ac *app.AppContainer) *gin.Engine {
	r.Use(middleware.Cors())

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/products", ac.ProductHandler.GetAllProducts)
		apiGroup.POST("/product-battle/submit", ac.BattleHandler.SubmitProductBattle)

		apiGroup.GET("/widgets/:widgetName", ac.WidgetHandler.Show)
	}

	return r
}
