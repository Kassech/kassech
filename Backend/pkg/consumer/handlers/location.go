package handlers

import (
	"encoding/json"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/workers"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

var bufferLock sync.Mutex

func HandleLocationMessage(msg amqp.Delivery) {
	var data struct {
		VehicleID uint    `json:"vehicle_id"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		PathID    uint    `json:"path_id,omitempty"`
	}

	log.Printf("Received message: %s", msg.Body)
	err := json.Unmarshal(msg.Body, &data)
	if err != nil {
		log.Printf("Error unmarshalling location message: %v", err)
		return
	}

	// Create properly timestamped log entry
	gpsLog := models.VehicleGPSLog{
		VehicleID: data.VehicleID,
		Latitude:  data.Lat,
		Longitude: data.Lon,
		PathID:    data.PathID,
		CreatedAt: time.Now().UTC(),
	}

	log.Printf("Received location message for vehicle %d, lat: %f, lon: %f",
		gpsLog.VehicleID, gpsLog.Latitude, gpsLog.Longitude)

	bufferLock.Lock()
	workers.LocationBuffer = append(workers.LocationBuffer, gpsLog)
	bufferLock.Unlock()
}
