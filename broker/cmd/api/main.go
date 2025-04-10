package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oscarxxi/product-battle/broker/internal/app"
	"github.com/oscarxxi/product-battle/broker/internal/app/configs"
	http_interface "github.com/oscarxxi/product-battle/broker/internal/interfaces/http"
)

func main() {
	configs.LoadConfig()

	appContainer := app.NewAppContainer()
	defer appContainer.Close()

	router := gin.Default()
	router = http_interface.SetApiRoutes(router, appContainer.ProductHandler)

	router.Run(configs.App.URL)
}
