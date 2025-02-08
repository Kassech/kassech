package handlers

import (
	"encoding/json"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/workers"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

var bufferLock sync.Mutex

// HandleLocationMessage processes incoming location messages
func HandleLocationMessage(msg amqp.Delivery) {
	var data models.VehicleGPSLog
	err := json.Unmarshal(msg.Body, &data)
	if err != nil {
		log.Println("Error parsing location JSON:", err)
		return
	}
	// Add to batch for PostgreSQL insert
	bufferLock.Lock()
	workers.LocationBuffer = append(workers.LocationBuffer, data)
	log.Printf("Added to LocationBuffer, current size: %d", len(workers.LocationBuffer))
	bufferLock.Unlock()
}
