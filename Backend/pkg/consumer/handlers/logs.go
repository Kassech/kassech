package handlers

import (
	"log"

	"github.com/streadway/amqp"
)

// HandleLogEvent processes log messages
func HandleLogEvent(msg amqp.Delivery) {
	log.Printf("Received Log: %s", msg.Body)
}
