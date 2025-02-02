package config

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection
var RabbitMQChannel *amqp.Channel

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

	// Open a channel
	RabbitMQChannel, err = RabbitMQConn.Channel()
	if err != nil {
		log.Fatal("Failed to open RabbitMQ channel:", err)
	}

	// Declare a queue
	queueName := os.Getenv("RABBITMQ_QUEUE") // Default: "location_updates"
	if queueName == "" {
		queueName = "location_updates" // Default queue name
	}

	_, err = RabbitMQChannel.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare queue:", err)
	}

	log.Println("RabbitMQ initialized")
}
