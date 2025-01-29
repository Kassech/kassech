package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"kassech/backend/pkg/websocket/middleware"
	"kassech/backend/pkg/websocket/server"
	"kassech/backend/pkg/websocket/service"

	"github.com/gorilla/websocket"
)

type PassengerHandler struct {
	connManager      *server.ConnectionManager
	passengerService *service.PassengerService
	auth             *middleware.WebSocketAuth
}

func NewPassengerHandler(
	connManager *server.ConnectionManager,
	passengerService *service.PassengerService,
	auth *middleware.WebSocketAuth,
) *PassengerHandler {
	return &PassengerHandler{
		connManager:      connManager,
		passengerService: passengerService,
		auth:             auth,
	}
}

func (h *PassengerHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	// Authenticate user
	userID, err := h.auth.Authenticate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Upgrade connection
	upgrader := &websocket.Upgrader{}        // Create a pointer to websocket.Upgrader
	conn, err := upgrader.Upgrade(w, r, nil) // Call Upgrade on the pointer
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer h.cleanupConnection(userID, conn)

	// Register connection
	h.connManager.AddConnection(userID, conn)

	// Maintain connection
	h.listenForMessages(userID, conn)
}

func (h *PassengerHandler) cleanupConnection(userID uint, conn *websocket.Conn) {
	conn.Close()
	h.connManager.RemoveConnection(userID)
}

func (h *PassengerHandler) listenForMessages(userID uint, conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Unexpected close: %v", err)
			}
			break
		}

		var msg map[string]interface{}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		switch msg["action"] {
		case "increment":
			pathID := uint(msg["pathID"].(float64))
			amount := int(msg["amount"].(float64))
			if err := h.passengerService.IncrementPassengerCountBy(pathID, amount); err != nil {
				log.Printf("Error incrementing passenger count: %v", err)
			}
			h.broadcastPassengerCount(pathID)
		case "decrement":
			pathID := uint(msg["pathID"].(float64))
			amount := int(msg["amount"].(float64))
			if err := h.passengerService.DecrementPassengerCountBy(pathID, amount); err != nil {
				log.Printf("Error decrementing passenger count: %v", err)
			}
			h.broadcastPassengerCount(pathID)
		case "getPassengers":
			pathID := uint(msg["pathID"].(float64))
			h.broadcastPassengerCount(pathID)
		default:
			log.Printf("Unknown action: %s", msg["action"])
		}
	}
}

func (h *PassengerHandler) broadcastPassengerCount(pathID uint) {
	// Fetch the current passenger count from the service
	passengerCount, err := h.passengerService.GetPassengerCount(pathID)
	if err != nil {
		log.Printf("Error fetching passenger count: %v", err)
		return
	}

	// Broadcast the passenger count to all connected clients
	message := map[string]interface{}{
		"pathID":         pathID,
		"passengerCount": passengerCount,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return
	}
	fmt.Println("	fmt", "fmt")
	h.connManager.Broadcast(messageBytes)
}
