package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"kassech/backend/pkg/config"
	"kassech/backend/pkg/websocket/middleware"
	"kassech/backend/pkg/websocket/service"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type LocationHandler struct {
	connManager     *config.ConnectionManager
	locationService *service.LocationService
	auth            *middleware.WebSocketAuth
	mu              sync.Mutex
}

type connectionState struct {
	writeChan chan []byte
	closeChan chan struct{}
	wg        sync.WaitGroup
}

type ClientMessage struct {
	Action    string   `json:"action"`
	Type      string   `json:"type,omitempty"`
	VehicleID *uint    `json:"vehicle_id,omitempty"`
	PathID    *uint    `json:"path_id,omitempty"`
	Lat       *float64 `json:"lat,omitempty"`
	Lon       *float64 `json:"lon,omitempty"`
	Radius    *float64 `json:"radius,omitempty"`
	Update    *struct {
		VehicleID uint    `json:"vehicle_id"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		PathID    *uint   `json:"path_id"`
	} `json:"update,omitempty"`
}

func NewLocationHandler(
	connManager *config.ConnectionManager,
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
	log.Println("Handling connection")

	userID, err := h.auth.Authenticate(r)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	log.Printf("Authenticated user: %d", userID)

	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	log.Println("Upgraded connection to WebSocket")

	state := &connectionState{
		writeChan: make(chan []byte, 100),
		closeChan: make(chan struct{}),
	}

	h.connManager.AddConnection(userID, conn)
	log.Printf("Added connection to connection manager for user: %d", userID)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		h.cleanupConnection(userID, conn, state)
		log.Printf("Cleaned up connection for user: %d", userID)
	}()

	conn.SetPingHandler(func(appData string) error {
		return conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(10*time.Second))
	})

	state.wg.Add(2)
	go h.handleWrites(conn, state)
	go h.handleReads(conn, state)
	log.Println("Started goroutines for handling reads and writes")

	h.listenForMessages(ctx, conn, state)
}

func (h *LocationHandler) handleReads(conn *websocket.Conn, state *connectionState) {
	defer state.wg.Done()

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	for {
		select {
		case <-state.closeChan:
			return
		default:
			_, _, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("Unexpected close: %v", err)
				}
				return
			}
			conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		}
	}
}

func (h *LocationHandler) handleWrites(conn *websocket.Conn, state *connectionState) {
	defer state.wg.Done()

	for {
		select {
		case message, ok := <-state.writeChan:
			if !ok {
				return
			}
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error writing to WebSocket: %v", err)
				return
			}
		case <-state.closeChan:
			return
		}
	}
}

func (h *LocationHandler) cleanupConnection(userID uint, conn *websocket.Conn, state *connectionState) {
	h.mu.Lock()
	defer h.mu.Unlock()

	close(state.closeChan)
	state.wg.Wait()
	conn.Close()
	h.connManager.RemoveConnection(userID, conn)
	close(state.writeChan)
}

func (h *LocationHandler) listenForMessages(ctx context.Context, conn *websocket.Conn, state *connectionState) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in listenForMessages: %v", r)
		}
	}()

	for {
		select {
		case <-state.closeChan:
			log.Println("Closing listenForMessages due to client close")
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("Unexpected close: %v", err)
				}
				log.Println("Closing listenForMessages due to error")
				return
			}

			log.Println("Received message")

			var msg ClientMessage
			if err := json.Unmarshal(message, &msg); err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			log.Printf("Handling message of type: %s", msg.Action)

			h.handleMessage(ctx, msg, state)
		}
	}
}

func (h *LocationHandler) handleMessage(ctx context.Context, msg ClientMessage, state *connectionState) {
	if msg.Action == "subscribe" {
		h.handleSubscription(ctx, msg, state)
		return
	}

	var locationMsg struct {
		VehicleID uint    `json:"vehicle_id"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		PathID    *uint   `json:"path_id"`
	}

	if msg.Update != nil {
		locationMsg = *msg.Update
	} else {
		if msg.VehicleID == nil || msg.Lat == nil || msg.Lon == nil {
			log.Printf("Missing required fields in location update")
			return
		}
		locationMsg = struct {
			VehicleID uint    `json:"vehicle_id"`
			Lat       float64 `json:"lat"`
			Lon       float64 `json:"lon"`
			PathID    *uint   `json:"path_id"`
		}{
			VehicleID: *msg.VehicleID,
			Lat:       *msg.Lat,
			Lon:       *msg.Lon,
			PathID:    msg.PathID,
		}
	}

	log.Printf("Handling location update for vehicle %d, lat: %f, lon: %f", locationMsg.VehicleID, locationMsg.Lat, locationMsg.Lon)
	h.processLocationUpdate(ctx, locationMsg, state)
}

func (h *LocationHandler) handleSubscription(ctx context.Context, msg ClientMessage, state *connectionState) {
	switch msg.Type {
	case "vehicle":
		if msg.VehicleID != nil {
			go h.subscribeToVehicle(ctx, *msg.VehicleID, state)
		}
	case "path":
		if msg.PathID != nil {
			go h.subscribeToPath(ctx, *msg.PathID, state)
		}
	case "all":
		go h.subscribeToAll(ctx, state)
	case "nearby":
		if msg.Lat != nil && msg.Lon != nil && msg.Radius != nil {
			go h.subscribeToNearby(ctx, *msg.Lat, *msg.Lon, *msg.Radius, state)
		}
	}
}

func (h *LocationHandler) processLocationUpdate(ctx context.Context, locationMsg struct {
	VehicleID uint    `json:"vehicle_id"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	PathID    *uint   `json:"path_id"`
}, _ *connectionState) {
	locationData, _ := json.Marshal(locationMsg)
	redisKey := fmt.Sprintf("vehicle_location:%d", locationMsg.VehicleID)

	log.Printf("Saving location to Redis for vehicle %d", locationMsg.VehicleID)
	if err := config.RedisClient.Set(ctx, redisKey, locationData, 0).Err(); err != nil {
		log.Printf("Error saving location to Redis: %v", err)
	}

	log.Printf("Updating GEO set for vehicle %d", locationMsg.VehicleID)
	if err := config.RedisClient.GeoAdd(
		ctx,
		"vehicle_locations",
		&redis.GeoLocation{
			Name:      strconv.Itoa(int(locationMsg.VehicleID)),
			Longitude: locationMsg.Lon,
			Latitude:  locationMsg.Lat,
		},
	).Err(); err != nil {
		log.Printf("Error updating GEO set: %v", err)
	}

	log.Printf("Publishing location update to all_vehicles and vehicle:%d channels", locationMsg.VehicleID)
	err := config.EventEmitter.Emit("location_updates", locationMsg)

	if err != nil {
		log.Printf("Failed to emit log event: %v", err)
	}

	allVehiclesChannel := "all_vehicles"
	vehicleChannel := fmt.Sprintf("vehicle:%d", locationMsg.VehicleID)
	config.RedisClient.Publish(ctx, vehicleChannel, locationData)
	config.RedisClient.Publish(ctx, allVehiclesChannel, locationData)

	if locationMsg.PathID != nil {
		pathChannel := fmt.Sprintf("path:%d", *locationMsg.PathID)
		log.Printf("Publishing location update to path:%d channel", *locationMsg.PathID)
		config.RedisClient.Publish(ctx, pathChannel, locationData)
	}
}

func (h *LocationHandler) subscribeToVehicle(ctx context.Context, vehicleID uint, state *connectionState) {
	state.wg.Add(1)
	defer state.wg.Done()

	channel := fmt.Sprintf("vehicle:%d", vehicleID)
	pubsub := config.RedisClient.Subscribe(ctx, channel)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for {
		select {
		case msg := <-ch:
			select {
			case state.writeChan <- []byte(msg.Payload):
			case <-state.closeChan:
				return
			}
		case <-state.closeChan:
			return
		}
	}
}

func (h *LocationHandler) subscribeToPath(ctx context.Context, pathID uint, state *connectionState) {
	state.wg.Add(1)
	defer state.wg.Done()

	channel := fmt.Sprintf("path:%d", pathID)
	pubsub := config.RedisClient.Subscribe(ctx, channel)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for {
		select {
		case msg := <-ch:
			select {
			case state.writeChan <- []byte(msg.Payload):
			case <-state.closeChan:
				return
			}
		case <-state.closeChan:
			return
		}
	}
}

func (h *LocationHandler) subscribeToAll(ctx context.Context, state *connectionState) {
	log.Println("Subscribing to all vehicles")
	state.wg.Add(1)
	defer state.wg.Done()

	pubsub := config.RedisClient.Subscribe(ctx, "all_vehicles")
	defer func() {
		log.Println("Closing all vehicles subscription")
		pubsub.Close()
	}()

	ch := pubsub.Channel()
	for {
		select {
		case msg := <-ch:
			log.Println("Received message from all vehicles subscription")
			select {
			case state.writeChan <- []byte(msg.Payload):
				log.Println("Sent message to client")
			case <-state.closeChan:
				log.Println("Closing subscription due to client close")
				return
			}
		case <-state.closeChan:
			log.Println("Closing subscription due to client close")
			return
		}
	}
}

func (h *LocationHandler) subscribeToNearby(ctx context.Context, lat, lon, radius float64, state *connectionState) {
	vehicleIDs, err := h.getNearbyVehicles(ctx, lat, lon, radius)
	if err != nil {
		log.Printf("Error getting nearby vehicles: %v", err)
		return
	}

	for _, vid := range vehicleIDs {
		state.wg.Add(1)
		go func(vehicleID uint) {
			defer state.wg.Done()
			h.subscribeToVehicle(ctx, vehicleID, state)
		}(vid)
	}
}

func (h *LocationHandler) getNearbyVehicles(ctx context.Context, lat, lon, radius float64) ([]uint, error) {
	query := &redis.GeoRadiusQuery{
		Radius: radius,
		Unit:   "km",
		Count:  100,
		Sort:   "ASC",
	}

	results, err := config.RedisClient.GeoRadius(ctx, "vehicle_locations", lon, lat, query).Result()
	if err != nil {
		return nil, err
	}

	var vehicleIDs []uint
	for _, loc := range results {
		vid, err := strconv.Atoi(loc.Name)
		if err != nil {
			continue
		}
		vehicleIDs = append(vehicleIDs, uint(vid))
	}
	return vehicleIDs, nil
}
