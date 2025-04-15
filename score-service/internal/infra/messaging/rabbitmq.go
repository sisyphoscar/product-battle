package messaging

import (
	"log"

	"github.com/oscarxxi/product-battle/score-service/internal/configs"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQ initializes a new RabbitMQ connection
func NewRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial(configs.Queue.RabbitMQURL)
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
