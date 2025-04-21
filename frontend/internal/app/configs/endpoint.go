package configs

import (
	"log"
	"os"
)

type endpointConfig struct {
	Broker string
}

var Endpoint endpointConfig

func loadEndpointConfig() {
	Endpoint = endpointConfig{
		Broker: os.Getenv("BROKER_ENDPOINT"),
	}
	if Endpoint.Broker == "" {
		log.Fatal("BROKER_ENDPOINT is not set")
	}

	log.Println("Endpoint config loaded")
}
