// consumer/consumer.go
package consumer

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

// ConsumeQueue listens to a queue and reconnects if the channel is closed.
func ConsumeQueue(conn *amqp.Connection, queue string, handler func(amqp.Delivery)) {
	for {
		ch, err := conn.Channel()
		if err != nil {
			log.Printf("Failed to open channel: %v. Retrying in 5 seconds...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Monitor channel close notifications.
		notifyClose := ch.NotifyClose(make(chan *amqp.Error))
		msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
		if err != nil {
			log.Printf("Failed to consume queue %s: %v. Retrying in 5 seconds...", queue, err)
			ch.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("Listening to queue: %s", queue)

		// Process messages in a separate goroutine.
		done := make(chan bool)
		go func() {
			for msg := range msgs {
				handler(msg)
			}
			done <- true
		}()

		// Wait for a channel close or errors.
		select {
		case err := <-notifyClose:
			log.Printf("Channel closed: %v. Reconnecting...", err)
		case <-done:
			log.Printf("Consumer for queue %s ended unexpectedly. Reconnecting...", queue)
		}

		ch.Close()
		time.Sleep(5 * time.Second)
	}
}
