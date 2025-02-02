package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"kassech/backend/pkg/config"
	"kassech/backend/pkg/websocket/middleware"
	"kassech/backend/pkg/websocket/server"
	"kassech/backend/pkg/websocket/service"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

type LocationHandler struct {
	connManager     *server.ConnectionManager
	locationService *service.LocationService
	auth            *middleware.WebSocketAuth
}

func NewLocationHandler(
	connManager *server.ConnectionManager,
	locationService *service.LocationService,
	auth *middleware.WebSocketAuth,
) *LocationHandler {
	return &LocationHandler{
		connManager:     connManager,
		locationService: locationService,
		auth:            auth,
	}
}

func (h *LocationHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
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

func (h *LocationHandler) cleanupConnection(userID uint, conn *websocket.Conn) {
	conn.Close()
	h.connManager.RemoveConnection(userID, conn)
}

func (h *LocationHandler) listenForMessages(userID uint, conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Unexpected close: %v", err)
			}
			break
		}

		var locationMsg struct {
			VehicleID uint    `json:"vehicle_id"`
			Lat       float64 `json:"lat"`
			Lon       float64 `json:"lon"`
			PathID    uint    `json:"path_id"`
		}

		// Unmarshal the incoming JSON message into the locationMsg struct
		err = json.Unmarshal(message, &locationMsg)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		// Update the vehicle's location
		// body, _ := json.Marshal(message)

		err = config.RabbitMQChannel.Publish("", "location_updates", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})

		// err = h.locationService.UpdateLocation(locationMsg.VehicleID, locationMsg.Lat, locationMsg.Lon, locationMsg.PathID)
		if err != nil {
			log.Printf("Error updating location: %v", err)
		}
	}
}
