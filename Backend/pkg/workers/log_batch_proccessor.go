// workers/log_worker.go
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
	LogBuffer         []models.SystemLog
	logBufferLock     sync.Mutex
	flushIntervalLogs = 5 * time.Second
	batchSizeLogs     = 50
)

func FlushLogs() {
	logBufferLock.Lock()
	defer logBufferLock.Unlock()

	if len(LogBuffer) == 0 {
		return
	}

	var queryBuilder strings.Builder
	queryBuilder.WriteString("INSERT INTO system_logs (event_type, entity_type, entity_id, details, created_at) VALUES ")

	values := []interface{}{}
	placeholders := []string{}

	for i, logEvent := range LogBuffer {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", 5*i+1, 5*i+2, 5*i+3, 5*i+4, 5*i+5))
		values = append(values, logEvent.EventType, logEvent.EntityType, logEvent.EntityID, logEvent.Details, logEvent.CreatedAt)
	}

	queryBuilder.WriteString(strings.Join(placeholders, ", "))

	_, err := database.DB.ConnPool.ExecContext(context.Background(), queryBuilder.String(), values...)
	if err != nil {
		log.Println("Batch Insert Error for logs:", err)
	} else {
		log.Printf("Inserted %d logs", len(LogBuffer))
	}

	LogBuffer = nil
}

func StartLogBatchFlusher() {
	go func() {
		for {
			time.Sleep(flushIntervalLogs)
			FlushLogs()
		}
	}()
}
