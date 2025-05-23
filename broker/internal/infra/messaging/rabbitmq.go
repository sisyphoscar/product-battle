package messaging

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sisyphoscar/product-battle/broker/internal/app/configs"
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
		log.Printf("Failed to connect to RabbitMQ (attempt %d/%d): %v", i+1, MAX_RETRIES, err)
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
func (r *RabbitMQ) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	return r.conn.Close()
}

// Create channel if the current channel is closed
func newChannelIfClosed(r *RabbitMQ) error {
	if r.channel.IsClosed() {
		newChannel, err := r.conn.Channel()
		if err != nil {
			return err
		}
		r.channel = newChannel
	}
	return nil
}
