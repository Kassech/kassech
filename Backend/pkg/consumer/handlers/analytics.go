package handlers

import (
	"log"

	"github.com/streadway/amqp"
)

// HandleAnalyticsEvent processes analytics messages
func HandleAnalyticsEvent(msg amqp.Delivery) {
	log.Printf("Processing Analytics Data: %s", msg.Body)
}
