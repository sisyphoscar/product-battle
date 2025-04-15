package messaging

import (
	"encoding/json"
	"log"
)

type Consumer interface {
	Listen(queue string, handler func([]byte) error) error
}

type ConsumerImpl struct {
	rabbitMQ *RabbitMQ
}

// NewConsumer creates a new RabbitMQ consumer
func NewConsumer(rabbitMQ *RabbitMQ) *ConsumerImpl {
	return &ConsumerImpl{
		rabbitMQ: rabbitMQ,
	}
}

/*
Listen starts consuming messages from the specified queue.
TODO: inject the handler function to process messages.
*/
func (c *ConsumerImpl) Listen(queue string) error {
	messages, err := c.rabbitMQ.channel.Consume(
		queue,
		"",    // consumer tag
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range messages {
			log.Printf("Received message: %s", msg.Body)
			err := dummyHandle(msg.Body)
			if err != nil {
				log.Printf("Error handling message: %v", err)
				msg.Nack(false, true) // requeue the message
			} else {
				msg.Ack(false) // acknowledge the message
			}
		}
	}()

	return nil
}

type SubmitProductBattleRequest struct {
	SeasonID      string         `json:"seasonId" binding:"required"`
	BattleResults []BattleResult `json:"roundResults" binding:"required,dive"`
}

type BattleResult struct {
	Round    int    `json:"round" binding:"required,gt=0"`
	WinnerID string `json:"winnerId" binding:"required"`
	LoserID  string `json:"loserId" binding:"required"`
}

func dummyHandle(msg []byte) error {
	var request SubmitProductBattleRequest
	err := json.Unmarshal(msg, &request)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return err
	}

	// Process the message here...

	return nil
}
