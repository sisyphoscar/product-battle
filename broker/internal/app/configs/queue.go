package configs

import (
	"log"
	"os"
)

type QueueConfig struct {
	RabbitMQURL       string
	BattleResultQueue string
}

var Queue QueueConfig

func loadQueueConfig() {
	Queue = QueueConfig{
		RabbitMQURL:       os.Getenv("RABBIT_MQ_DSN"),
		BattleResultQueue: os.Getenv("BATTLE_RESULT_QUEUE"),
	}
	if Queue.RabbitMQURL == "" {
		log.Fatal("RABBIT_MQ_DSN is not set")
	}
	if Queue.BattleResultQueue == "" {
		log.Fatal("BATTLE_RESULT_QUEUE is not set")
	}

	log.Println("Queue config loaded")
}
