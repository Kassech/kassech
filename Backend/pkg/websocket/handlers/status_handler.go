package handlers

import (
	"log"
	"net/http"

	"kassech/backend/pkg/config"
	"kassech/backend/pkg/websocket/middleware"
	"kassech/backend/pkg/websocket/service"

	"github.com/gorilla/websocket"
)

type StatusHandler struct {
	connManager   *config.ConnectionManager
	statusService *service.StatusService
	auth          *middleware.WebSocketAuth
}

func NewStatusHandler(
	connManager *config.ConnectionManager,
	statusService *service.StatusService,
	auth *middleware.WebSocketAuth,
) *StatusHandler {
	return &StatusHandler{
		connManager:   connManager,
		statusService: statusService,
		auth:          auth,
	}
}

func (h *StatusHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
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

	// Update status
	if err := h.statusService.SetOnline(userID); err != nil {
		log.Printf("Status update failed: %v", err)
	}

	// Maintain connection
	h.listenForMessages(userID, conn)
}

func (h *StatusHandler) cleanupConnection(userID uint, conn *websocket.Conn) {
	conn.Close()
	h.connManager.RemoveConnection(userID, conn)
	h.statusService.SetOffline(userID)
}

func (h *StatusHandler) listenForMessages(userID uint, conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Unexpected close: %v", err)
			}
			break
		}
	}
}
