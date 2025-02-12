// pkg/event/emitter.go
package event

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type EventEmitter struct {
	channel *amqp.Channel
}

func NewEventEmitter(channel *amqp.Channel) *EventEmitter {
	return &EventEmitter{channel: channel}
}

func (e *EventEmitter) Emit(eventType string, payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling payload: %v", err)
	}

	err = e.channel.Publish(
		"",        // exchange
		eventType, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
