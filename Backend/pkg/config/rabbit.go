package config

import (
	"kassech/backend/pkg/event"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection
var EventEmitter *event.EventEmitter

// InitRabbitMQ initializes the RabbitMQ connection and declares the queue.
func InitRabbitMQ() {
	// Retrieve RabbitMQ connection details from environment variables
	rabbitMQURL := os.Getenv("RABBITMQ_URL") // Default: "amqp://guest:guest@localhost:5672/"
	if rabbitMQURL == "" {
		rabbitMQURL = "amqp://guest:guest@localhost:5672/" // Fallback to default if not set
	}

	// Connect to RabbitMQ
	var err error
	RabbitMQConn, err = amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	log.Println("RabbitMQ initialized")

	conn, err := RabbitMQConn.Channel()
	if err != nil {
		log.Fatal("Failed to open RabbitMQ channel:", err)
	}
	EventEmitter = event.NewEventEmitter(conn)
}
