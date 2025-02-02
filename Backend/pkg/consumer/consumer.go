package consumer

import (
	"kassech/backend/pkg/consumer/handlers"
	"log"

	"github.com/streadway/amqp"
)

// MessageHandler defines a function type for handling messages.
type MessageHandler func(amqp.Delivery)

// Map of queue names to handler functions.
var handler = map[string]MessageHandler{
	"location_updates": handlers.HandleLocationMessage,
	"logs":             handlers.HandleLogEvent,
	"analytics":        handlers.HandleAnalyticsEvent,
}

// ConsumeQueue listens to a queue and processes messages with the appropriate handler.
func ConsumeQueue(ch *amqp.Channel, queue string) {
	msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to consume queue %s: %v", queue, err)
	}

	handler, exists := handler[queue]
	if !exists {
		log.Fatalf("No handler registered for queue: %s", queue)
	}

	log.Printf("Listening to queue: %s", queue)
	for msg := range msgs {
		handler(msg)
	}
}
