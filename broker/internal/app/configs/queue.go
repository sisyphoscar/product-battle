package configs

import (
	"log"
	"os"
)

type QueueConfig struct {
	RabbitMQURL          string
	ScoreSettlementQueue string
}

var Queue QueueConfig

func loadQueueConfig() {
	Queue = QueueConfig{
		RabbitMQURL:          os.Getenv("RABBIT_MQ_DSN"),
		ScoreSettlementQueue: os.Getenv("SCORE_SETTLEMENT_QUEUE"),
	}
	if Queue.RabbitMQURL == "" {
		log.Fatal("RABBIT_MQ_DSN is not set")
	}
	if Queue.ScoreSettlementQueue == "" {
		log.Fatal("SCORE_SETTLEMENT_QUEUE is not set")
	}

	log.Println("Queue config loaded")
}
