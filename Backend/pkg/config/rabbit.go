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
func PublishToQueue(queueName string, body []byte) error {
	ch, err := RabbitMQConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName,
		false, // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",        // Exchange
		queueName, // Routing key
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
