package workers

import (
	"context"
	"fmt"
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	LocationBuffer []models.VehicleGPSLog
	flushInterval  = 3 * time.Second
	bufferLock     sync.Mutex
	batchSize      = 1000
)

// FlushLocations inserts buffered locations into PostgreSQL
func FlushLocations() {
	bufferLock.Lock()
	defer bufferLock.Unlock()

	if len(LocationBuffer) == 0 {
		return
	}
	// Build SQL query
	var queryBuilder strings.Builder
	queryBuilder.WriteString("INSERT INTO vehicle_gps_logs (vehicle_id, latitude, longitude, path_id, created_at) VALUES ")

	values := []interface{}{}
	placeholders := []string{}

	for i, loc := range LocationBuffer {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", 5*i+1, 5*i+2, 5*i+3, 5*i+4, 5*i+5))
		values = append(values, loc.VehicleID, loc.Latitude, loc.Longitude, loc.PathID, loc.CreatedAt)
	}

	queryBuilder.WriteString(strings.Join(placeholders, ", "))

	// Execute batch insert
	_, err := database.DB.ConnPool.ExecContext(context.Background(), queryBuilder.String(), values...)
	if err != nil {
		log.Println("Batch Insert Error:", err)
	} else {
		log.Printf("Inserted %d locations", len(LocationBuffer))
	}

	// Clear buffer
	LocationBuffer = nil
}

// StartBatchFlusher runs a background task to flush locations every X seconds
func StartBatchFlusher() {
	go func() {
		for {
			time.Sleep(flushInterval)
			FlushLocations()
		}
	}()
}
