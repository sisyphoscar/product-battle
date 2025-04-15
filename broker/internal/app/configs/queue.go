package configs

import (
	"log"
	"os"
)

type QueueConfig struct {
	RabbitMQURL      string
	BattleScoreQueue string
}

var Queue QueueConfig

func loadQueueConfig() {
	Queue = QueueConfig{
		RabbitMQURL:      os.Getenv("RABBIT_MQ_DSN"),
		BattleScoreQueue: os.Getenv("BATTLE_SCORE_QUEUE"),
	}
	if Queue.RabbitMQURL == "" {
		log.Fatal("RABBIT_MQ_DSN is not set")
	}
	if Queue.BattleScoreQueue == "" {
		log.Fatal("BATTLE_SCORE_QUEUE is not set")
	}

	log.Println("Queue config loaded")
}
