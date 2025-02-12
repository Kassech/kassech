package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"sync"
	"time"

	"kassech/backend/pkg/config"
	"kassech/backend/pkg/websocket/middleware"
	"kassech/backend/pkg/websocket/service"

	"github.com/gorilla/websocket"
)

type PassengerHandler struct {
	connManager      *config.ConnectionManager
	passengerService *service.PassengerService
	auth             *middleware.WebSocketAuth
}

type passengerConnectionState struct {
	closeChan chan struct{}
	wg        sync.WaitGroup
	userID    uint
}

func NewPassengerHandler(
	connManager *config.ConnectionManager,
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
	userID, err := h.auth.Authenticate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	state := &passengerConnectionState{
		closeChan: make(chan struct{}),
		userID:    userID,
	}

	h.connManager.AddConnection(userID, conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		h.cleanupConnection(userID, conn, state)
	}()

	conn.SetPingHandler(func(appData string) error {
		return conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(10*time.Second))
	})

	state.wg.Add(1)
	go h.handleReads(conn, state)

	h.handleConnectionMessages(ctx, conn, state)
}

func (h *PassengerHandler) handleReads(conn *websocket.Conn, state *passengerConnectionState) {
	defer state.wg.Done()

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	for {
		select {
		case <-state.closeChan:
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("Unexpected close: %v", err)
				}
				return
			}

			var msg map[string]interface{}
			if err := json.Unmarshal(message, &msg); err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			h.handleMessage(msg, state)
		}
	}
}

func (h *PassengerHandler) cleanupConnection(userID uint, conn *websocket.Conn, state *passengerConnectionState) {
	close(state.closeChan)
	state.wg.Wait()
	conn.Close()
	h.connManager.RemoveConnection(userID, conn)
}

func (h *PassengerHandler) handleConnectionMessages(_ context.Context, _ *websocket.Conn, state *passengerConnectionState) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n%s", r, debug.Stack())
		}
	}()

	for {
		select {
		case <-state.closeChan:
			log.Println("Connection closed, stopping message handler")
			return
		}
	}
}

func (h *PassengerHandler) handleMessage(msg map[string]interface{}, state *passengerConnectionState) {
	switch msg["action"] {
	case "increment":
		h.handleIncrement(msg, state)
	case "decrement":
		h.handleDecrement(msg, state)
	case "getPassengers":
		h.handleGetPassengers(msg, state)
	default:
		log.Printf("Unknown action: %s", msg["action"])
	}
}

func (h *PassengerHandler) sendPassengerCount(userID uint, pathID uint) {
	count, err := h.passengerService.GetPassengerCount(pathID)
	if err != nil {
		log.Printf("Error getting passenger count: %v", err)
		return
	}

	response := map[string]interface{}{
		"pathID":         pathID,
		"passengerCount": count,
	}

	message, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		return
	}

	h.connManager.SendToUser(userID, message)
}

func (h *PassengerHandler) handleIncrement(msg map[string]interface{}, state *passengerConnectionState) {
	pathID := uint(msg["pathID"].(float64))
	amount := int(msg["amount"].(float64))
	log.Printf("Incrementing passengers for path %d by %d", pathID, amount)

	if err := h.passengerService.IncrementPassengerCountBy(pathID, amount); err != nil {
		log.Printf("Error incrementing: %v", err)
		return
	}

	h.sendPassengerCount(state.userID, pathID)
	h.broadcastPassengerCount(pathID)
}

func (h *PassengerHandler) handleDecrement(msg map[string]interface{}, state *passengerConnectionState) {
	pathID := uint(msg["pathID"].(float64))
	amount := int(msg["amount"].(float64))
	log.Printf("Decrementing passengers for path %d by %d", pathID, amount)

	if err := h.passengerService.DecrementPassengerCountBy(pathID, amount); err != nil {
		log.Printf("Error decrementing: %v", err)
		return
	}

	h.sendPassengerCount(state.userID, pathID)
	h.broadcastPassengerCount(pathID)
}

func (h *PassengerHandler) handleGetPassengers(msg map[string]interface{}, state *passengerConnectionState) {
	pathID := uint(msg["pathID"].(float64))
	h.sendPassengerCount(state.userID, pathID)
}

func (h *PassengerHandler) broadcastPassengerCount(pathID uint) {
	count, err := h.passengerService.GetPassengerCount(pathID)
	if err != nil {
		log.Printf("Error getting count: %v", err)
		return
	}

	message := map[string]interface{}{
		"pathID":         pathID,
		"passengerCount": count,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling: %v", err)
		return
	}

	channel := fmt.Sprintf("passenger_updates:%d", pathID)
	config.RedisClient.Publish(context.Background(), channel, messageBytes)
}
