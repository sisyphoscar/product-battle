package messaging

import (
	"log"

	"github.com/oscarxxi/product-battle/score-service/internal/configs"
	"github.com/oscarxxi/product-battle/score-service/internal/domain/score"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ScoreConsumer struct {
	rabbitMQ     *RabbitMQ
	scoreService *score.ScoreService
}

// NewConsumer creates a new RabbitMQ consumer
func NewScoreConsumer(rabbitMQ *RabbitMQ, scoreService *score.ScoreService) *ScoreConsumer {
	return &ScoreConsumer{
		rabbitMQ:     rabbitMQ,
		scoreService: scoreService,
	}
}

// Listen starts consuming messages from the specified queue.
func (c *ScoreConsumer) Listen() error {
	queueName := configs.Queue.BattleScoreQueue
	err := c.rabbitMQ.DeclareQueue(queueName)
	if err != nil {
		return err
	}

	log.Println("Listening for messages on queue:", queueName)

	messages, err := c.consumeFromQueue(queueName)
	if err != nil {
		return err
	}

	go c.handle(messages)

	return nil
}

// consume sets up the consumer to listen for messages on the queue.
func (c *ScoreConsumer) consumeFromQueue(queueName string) (<-chan amqp.Delivery, error) {
	messages, err := c.rabbitMQ.channel.Consume(
		queueName,
		"",    // consumer tag
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// handle processes incoming messages.
func (c *ScoreConsumer) handle(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		log.Printf("Received message: %s", msg.Body)
		err := c.scoreService.HandleBattleResults(msg.Body)
		if err != nil {
			log.Printf("Error handling message: %v", err)
			msg.Nack(false, false) // not requeue
		} else {
			log.Printf("Message processed successfully: %s", msg.Body)
			msg.Ack(false)
		}
	}
}
