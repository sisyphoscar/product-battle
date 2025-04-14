package messaging

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Publish publishes a message to a RabbitMQ queue
func (r *RabbitMQ) Publish(queueName string, body string) error {
	err := newChannelIfClosed(r)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Printf("Failed to publish message to queue %s: %v", queueName, err)
		return err
	}

	log.Printf("Message published to queue %s: %s", queueName, body)
	return nil
}
