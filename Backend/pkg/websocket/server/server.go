package server

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	connections map[uint]*websocket.Conn
	mu          sync.RWMutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[uint]*websocket.Conn),
	}
}

func (cm *ConnectionManager) AddConnection(userID uint, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.connections[userID] = conn
}

func (cm *ConnectionManager) RemoveConnection(userID uint) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.connections, userID)
}

func (cm *ConnectionManager) Broadcast(message []byte) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	fmt.Println("Current connections:", cm.connections)
	for userID, conn := range cm.connections {
		fmt.Printf("[DEBUG] Failed to send message to user %d:", userID)

		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Printf("[DEBUG] Failed to send message to user %d: %v\n", userID, err)
			// Optionally, remove the connection if it's no longer valid
			cm.RemoveConnection(userID)
		}
	}
}
