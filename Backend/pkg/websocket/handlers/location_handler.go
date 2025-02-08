package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"kassech/backend/pkg/config"
	"kassech/backend/pkg/websocket/middleware"
	"kassech/backend/pkg/websocket/server"
	"kassech/backend/pkg/websocket/service"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type LocationHandler struct {
	connManager     *server.ConnectionManager
	locationService *service.LocationService
	auth            *middleware.WebSocketAuth
	writeChan       chan []byte // Channel for serializing WebSocket writes
	closeOnce       sync.Once
}

type ClientMessage struct {
	Action    string   `json:"action"`
	Type      string   `json:"type,omitempty"`
	VehicleID *uint    `json:"vehicle_id,omitempty"`
	PathID    *uint    `json:"path_id,omitempty"`
	Lat       *float64 `json:"lat,omitempty"`
	Lon       *float64 `json:"lon,omitempty"`
	Radius    *float64 `json:"radius,omitempty"`
	// For location update without subscribe action
	Update *struct {
		VehicleID uint    `json:"vehicle_id"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		PathID    *uint   `json:"path_id"`
	} `json:"update,omitempty"`
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
		writeChan:       make(chan []byte, 100), // Buffered channel to avoid blocking
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

func (h *LocationHandler) cleanupConnection(userID uint, conn *websocket.Conn) {
	conn.Close()
	h.connManager.RemoveConnection(userID, conn)
	h.closeOnce.Do(func() {
		close(h.writeChan)
	})
}
func (h *LocationHandler) handleWrites(conn *websocket.Conn) {
	for message := range h.writeChan {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error writing to WebSocket: %v", err)
			return
		}
	}
}

func (h *LocationHandler) listenForMessages(_ uint, conn *websocket.Conn) {
	ctx := context.Background()

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

		// If action is subscribe, launch listener subscription
		if msg.Action == "subscribe" {
			switch msg.Type {
			case "vehicle":
				if msg.VehicleID != nil {
					go h.subscribeToVehicle(ctx, *msg.VehicleID)
				}
			case "path":
				if msg.PathID != nil {
					go h.subscribeToPath(ctx, *msg.PathID)
				}
			case "all": // Add new case for "all" type
				go h.subscribeToAll(ctx)
			case "nearby":
				if msg.Lat != nil && msg.Lon != nil && msg.Radius != nil {
					go h.subscribeToNearby(ctx, *msg.Lat, *msg.Lon, *msg.Radius)
				}
			}
			continue
		}

		// Otherwise, treat as a location update
		var locationMsg struct {
			VehicleID uint    `json:"vehicle_id"`
			Lat       float64 `json:"lat"`
			Lon       float64 `json:"lon"`
			PathID    *uint   `json:"path_id"`
		}
		// Prefer update block if provided
		if msg.Update != nil {
			locationMsg.VehicleID = msg.Update.VehicleID
			locationMsg.Lat = msg.Update.Lat
			locationMsg.Lon = msg.Update.Lon
			locationMsg.PathID = msg.Update.PathID
		} else {
			// Fallback to top-level fields if available
			if msg.VehicleID == nil || msg.Lat == nil || msg.Lon == nil {
				log.Printf("Missing required fields in location update")
				continue
			}
			locationMsg.VehicleID = *msg.VehicleID
			locationMsg.Lat = *msg.Lat
			locationMsg.Lon = *msg.Lon
			locationMsg.PathID = msg.PathID
		}

		// Save location to Redis (active state)
		redisKey := fmt.Sprintf("vehicle_location:%d", locationMsg.VehicleID)
		locationData, _ := json.Marshal(locationMsg)
		err = config.RedisClient.Set(ctx, redisKey, locationData, 0).Err()
		if err != nil {
			log.Printf("Error saving location to Redis: %v", err)
		}

		// Add to GEO set for nearby queries
		err = config.RedisClient.GeoAdd(
			ctx,
			"vehicle_locations",
			&redis.GeoLocation{
				Name:      strconv.Itoa(int(locationMsg.VehicleID)), // Vehicle ID as string
				Longitude: locationMsg.Lon,                          // LONGITUDE first
				Latitude:  locationMsg.Lat,                          // LATITUDE second
			},
		).Err()
		if err != nil {
			log.Printf("Error updating GEO set: %v", err)
		}
		err = config.EventEmitter.Emit("location_updates", locationData)
		if err != nil {
			log.Printf("Error publishing location update: %v", err)
		}
		// Publish location update to Redis channels for real-time listeners
		allVehiclesChannel := "all_vehicles"
		vehicleChannel := "vehicle:" + strconv.Itoa(int(locationMsg.VehicleID))
		config.RedisClient.Publish(ctx, vehicleChannel, locationData)
		config.RedisClient.Publish(ctx, allVehiclesChannel, locationData) // New publish

		if locationMsg.PathID != nil {
			pathChannel := "path:" + strconv.Itoa(int(*locationMsg.PathID))
			config.RedisClient.Publish(ctx, pathChannel, locationData)
		}
	}
}

func (h *LocationHandler) subscribeToVehicle(ctx context.Context, vehicleID uint) {
	channel := "vehicle:" + strconv.Itoa(int(vehicleID))
	pubsub := config.RedisClient.Subscribe(ctx, channel)
	defer pubsub.Close()
	ch := pubsub.Channel()
	for msg := range ch {
		h.writeChan <- []byte(msg.Payload) // Send message to write channel
	}
}

func (h *LocationHandler) subscribeToPath(ctx context.Context, pathID uint) {
	channel := "path:" + strconv.Itoa(int(pathID))
	pubsub := config.RedisClient.Subscribe(ctx, channel)
	defer pubsub.Close()
	ch := pubsub.Channel()
	for msg := range ch {
		h.writeChan <- []byte(msg.Payload) // Send message to write channel
	}
}

func (h *LocationHandler) subscribeToNearby(ctx context.Context, lat, lon, radius float64) {
	// Get initial list of nearby vehicles via Redis GEO
	vehicleIDs, err := getNearbyVehicles(ctx, lat, lon, radius)
	if err != nil {
		log.Printf("Error getting nearby vehicles: %v", err)
		return
	}
	// Subscribe to each vehicle's channel in separate goroutines
	for _, vid := range vehicleIDs {
		go h.subscribeToVehicle(ctx, vid)
	}
}

func getNearbyVehicles(ctx context.Context, lat, lon, radius float64) ([]uint, error) {
	query := &redis.GeoRadiusQuery{
		Radius:    radius,
		Unit:      "km",
		WithCoord: false,
		WithDist:  false,
		Sort:      "ASC",
		Count:     100, // Adjust as needed
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
func (h *LocationHandler) subscribeToAll(ctx context.Context) {
	pubsub := config.RedisClient.Subscribe(ctx, "all_vehicles")
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		// Send message to write channel for all connected clients
		h.writeChan <- []byte(msg.Payload)
	}
}
