package configs

import (
	"log"
	"os"
)

type QueueConfig struct {
	RabbitMQURL string
}

var Queue QueueConfig

func loadQueueConfig() {
	Queue = QueueConfig{
		RabbitMQURL: os.Getenv("RABBIT_MQ_DSN"),
	}
	if Queue.RabbitMQURL == "" {
		log.Fatal("RABBIT_MQ_DSN is not set")
	}

	log.Println("Queue config loaded")
}
