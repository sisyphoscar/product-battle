package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/oscarxxi/product-battle/score-service/internal/configs"
	"github.com/oscarxxi/product-battle/score-service/internal/domain/score"
	"github.com/oscarxxi/product-battle/score-service/internal/infra/db"
	"github.com/oscarxxi/product-battle/score-service/internal/infra/messaging"
	repository "github.com/oscarxxi/product-battle/score-service/internal/infra/repositories/postgres"
)

func main() {
	configs.LoadConfig()

	rabbitMQ, err := messaging.NewRabbitMQ()
	if err != nil {
		panic(err)
	}
	defer rabbitMQ.Close()

	db, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	scoreRepository := repository.NewScoreRepository(db)

	scoreService := score.NewScoreService(scoreRepository)
	scoreConsumer := messaging.NewScoreConsumer(rabbitMQ, scoreService)

	err = scoreConsumer.Listen()
	if err != nil {
		log.Fatalf("Error listening to queue: %v", err)
	}

	// Graceful shutdown, and keep queue listening
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")
}
