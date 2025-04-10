package configs

import (
	"log"
	"os"
)

type endpointConfig struct {
	ProductService string
}

var Endpoint endpointConfig

func loadEndpointConfig() {
	Endpoint = endpointConfig{
		ProductService: os.Getenv("PRODUCT_SERVICE_ENDPOINT"),
	}
	if Endpoint.ProductService == "" {
		log.Fatal("PRODUCT_SERVICE_ENDPOINT is not set")
	}

	log.Println("Endpoint config loaded")
}
