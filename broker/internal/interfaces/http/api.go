package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/app"
	middleware "github.com/oscarxxi/product-battle/broker/internal/interfaces/http/middlewares"
)

// Api sets up the API routes for the application
func SetApiRoutes(router *gin.Engine, ac *app.AppContainer) *gin.Engine {
	router.Use(middleware.Cors())

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/products", ac.ProductHandler.GetAllProducts)
		apiGroup.POST("/product-battle/submit", ac.BattleHandler.SubmitProductBattle)

		apiGroup.GET("widgets/:widgetName", func(c *gin.Context) {
			widget := c.Param("widgetName")

			if widget != "product-score" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid widget: " + widget})
				return
			}

			res, err := http.Get("http://localhost:8085/widgets/product-score")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get widget"})
				return
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get widget"})
				return
			}

			type WidgetResponse struct {
				Name  string `json:"name"`
				Stats []struct {
					ProductName string `json:"productName"`
					Score       int    `json:"score"`
				} `json:"stats"`
			}

			var widgetResponse WidgetResponse
			if err := json.NewDecoder(res.Body).Decode(&widgetResponse); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode widget response"})
				return
			}
			c.JSON(http.StatusOK, widgetResponse)

		})
	}

	return router
}
