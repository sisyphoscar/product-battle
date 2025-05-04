package messaging

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sisyphoscar/product-battle/score-service/internal/configs"
)

const (
	MAX_RETRIES    = 5
	RETRY_INTERVAL = 5 * time.Second
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQ initializes a new RabbitMQ connection
func NewRabbitMQ() (*RabbitMQ, error) {
	var conn *amqp.Connection
	var err error

	for i := 0; i < MAX_RETRIES; i++ {
		conn, err = amqp.Dial(configs.Queue.RabbitMQURL)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to RabbitMQ (attempt %d/5): %v", i+1, err)
		time.Sleep(RETRY_INTERVAL)
	}

	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	log.Println("RabbitMQ connected")

	return &RabbitMQ{
		conn:    conn,
		channel: channel,
	}, nil
}

// Close closes the RabbitMQ connection
func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
	log.Println("RabbitMQ connection closed")
}

// DeclareQueue ensures the queue exists before consuming messages
func (r *RabbitMQ) DeclareQueue(queueName string) error {
	_, err := r.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue %s: %w", queueName, err)
	}

	return nil
}
