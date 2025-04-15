package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/oscarxxi/product-battle/score-service/internal/configs"
	"github.com/oscarxxi/product-battle/score-service/internal/infra/messaging"
)

func main() {
	configs.LoadConfig()

	rabbitMQ, err := messaging.NewRabbitMQ()
	if err != nil {
		panic(err)
	}
	defer rabbitMQ.Close()

	consumer := messaging.NewConsumer(rabbitMQ)

	err = consumer.Listen(configs.Queue.BattleScoreQueue)
	if err != nil {
		log.Fatalf("Error listening to queue: %v", err)
	}

	// Graceful shutdown, and keep queue listening
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")
}
