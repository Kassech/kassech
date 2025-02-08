// handlers/log_handler.go
package handlers

import (
	"encoding/json"
	"kassech/backend/pkg/database"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type SystemLog struct {
	EventType  string    `json:"event_type"`
	EntityType string    `json:"entity_type"`
	EntityID   uint      `json:"entity_id"`
	Details    string    `json:"details"`
	CreatedAt  time.Time `json:"created_at"`
}

var (
	logBuffer     []SystemLog
	logBufferLock sync.Mutex
	batchSize     = 50
	flushInterval = 5 * time.Second
)

func HandleLogEvent(msg amqp.Delivery) {
	var event SystemLog
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		log.Println("Error parsing log event JSON:", err)
		return
	}
	logBufferLock.Lock()
	logBuffer = append(logBuffer, event)
	log.Printf("Added log event to buffer, current size: %d", len(logBuffer))
	logBufferLock.Unlock()
}

func FlushLogs() {
	logBufferLock.Lock()
	defer logBufferLock.Unlock()
	if len(logBuffer) == 0 {
		return
	}
	if err := database.DB.Create(&logBuffer).Error; err != nil {
		log.Println("Error inserting logs:", err)
	} else {
		log.Printf("Inserted %d logs into system_logs", len(logBuffer))
	}
	logBuffer = nil // Reset buffer
}

func StartLogFlusher() {
	go func() {
		ticker := time.NewTicker(flushInterval)
		defer ticker.Stop()
		for range ticker.C {
			FlushLogs()
		}
	}()
}
