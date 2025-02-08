package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"kassech/backend/pkg/websocket/middleware"
	"kassech/backend/pkg/websocket/server"
	"kassech/backend/pkg/websocket/service"

	"github.com/gorilla/websocket"
)

type DestinationHandler struct {
	connManager        *server.ConnectionManager
	destinationService *service.DestinationService
	auth               *middleware.WebSocketAuth
	writeChan          chan []byte // Channel for serializing WebSocket writes
	closeOnce          sync.Once   // Ensures writeChan is closed only once

}

func NewDestinationHandler(
	connManager *server.ConnectionManager,
	destinationService *service.DestinationService,
	auth *middleware.WebSocketAuth,
) *DestinationHandler {
	return &DestinationHandler{
		connManager:        connManager,
		destinationService: destinationService,
		auth:               auth,
		writeChan:          make(chan []byte, 100), // Buffered channel to avoid blocking
	}
}

func (h *DestinationHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	// Authenticate user
	userID, err := h.auth.Authenticate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Upgrade connection
	upgrader := &websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer h.cleanupConnection(userID, conn)

	// Register connection
	h.connManager.AddConnection(userID, conn)

	// Start a goroutine to handle writes to the WebSocket connection
	go h.handleWrites(conn)

	// Maintain connection
	h.listenForMessages(userID, conn)
}

func (h *DestinationHandler) cleanupConnection(userID uint, conn *websocket.Conn) {
	conn.Close()
	h.connManager.RemoveConnection(userID, conn)
	h.closeOnce.Do(func() {
		close(h.writeChan) // Close the write channel
	})
}

func (h *DestinationHandler) handleWrites(conn *websocket.Conn) {
	for message := range h.writeChan {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error writing to WebSocket: %v", err)
			return
		}
	}
}
func (h *DestinationHandler) listenForMessages(userID uint, conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Unexpected close: %v", err)
			}
			break
		}

		var msg ClientMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		// Broadcast the message to all connected users with the same ID
		h.connManager.SendToUser(userID, message)

	}
}
