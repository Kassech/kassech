package server

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	connections map[uint][]*websocket.Conn
	mu          sync.RWMutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[uint][]*websocket.Conn),
	}
}

func (cm *ConnectionManager) AddConnection(userID uint, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.connections[userID] = append(cm.connections[userID], conn)
}

func (cm *ConnectionManager) RemoveConnection(userID uint, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	connections := cm.connections[userID]
	for i, c := range connections {
		if c == conn {
			cm.connections[userID] = append(connections[:i], connections[i+1:]...)
			c.Close()
			break
		}
	}
	if len(cm.connections[userID]) == 0 {
		delete(cm.connections, userID)
	}
}

func (cm *ConnectionManager) Broadcast(message []byte) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	fmt.Println("Current connections:", cm.connections)
	for userID, conns := range cm.connections {
		for _, conn := range conns {
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Printf("[DEBUG] Failed to send message to user %d: %v\n", userID, err)
				cm.RemoveConnection(userID, conn)
			}
		}
	}
}
