// consumer/consumer.go
package consumer

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

// ConsumeQueue listens to a queue with proper error handling and reconnection logic
func ConsumeQueue(conn *amqp.Connection, queue string, handler func(amqp.Delivery)) {
	for {
		ch, err := conn.Channel()
		if err != nil {
			log.Printf("Failed to open channel: %v. Retrying in 5 seconds...\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// 1. Declare queue first to ensure existence
		_, err = ch.QueueDeclare(
			queue,
			true,  // Durable
			false, // AutoDelete
			false, // Exclusive
			false, // NoWait
			nil,
		)
		if err != nil {
			log.Printf("Failed to declare queue: %v", err)
			ch.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		// 2. Set QoS to prevent overwhelming consumer
		err = ch.Qos(
			1,     // Prefetch count
			0,     // Prefetch size
			false, // Global
		)
		if err != nil {
			log.Printf("Failed to set QoS: %v", err)
			ch.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		// 3. Start consuming with manual acknowledgments
		msgs, err := ch.Consume(
			queue,
			"",    // Consumer tag
			false, // Auto-ack (MUST set to false)
			false, // Exclusive
			false, // NoLocal
			false, // NoWait
			nil,
		)
		if err != nil {
			log.Printf("Failed to consume queue %s: %v. Retrying...\n", queue, err)
			ch.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("Successfully started consuming queue: %s", queue)

		// 4. Process messages with recovery mechanism
		for msg := range msgs {
			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Recovered from panic in handler: %v", r)
					}
				}()

				// Process message
				handler(msg)

				// Manual acknowledgment after successful processing
				if err := msg.Ack(false); err != nil {
					log.Printf("Failed to acknowledge message: %v", err)
				}
			}()
		}

		// 5. Handle channel closure
		log.Printf("Message channel closed for queue %s. Reconnecting...", queue)
		ch.Close()
		time.Sleep(5 * time.Second)
	}
}
