package handlers

import (
	"context"
	"encoding/json"
	"kassech/backend/pkg/config"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/workers"
	"log"
	"strconv"
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

	// Store latest location in Redis (real-time tracking)
	ctx := context.Background()
	redisKey := "location:" + strconv.FormatUint(uint64(data.VehicleID), 10)
	config.RedisClient.HSet(ctx, redisKey, "lat", data.Latitude, "lon", data.Longitude)
	log.Printf("Stored in Redis with key: %s", redisKey)
}
