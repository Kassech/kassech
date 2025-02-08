package main

import (
	"fmt"
	"kassech/backend/pkg/config"
	"kassech/backend/pkg/consumer"
	"kassech/backend/pkg/consumer/handlers"
	"kassech/backend/pkg/controller"
	"kassech/backend/pkg/database"
	"kassech/backend/pkg/repository"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/websocket"
	"kassech/backend/pkg/workers"
	"log"
	"os"
	"time"

	routes "kassech/backend/pkg/delivery/http"

	scripts "kassech/backend/script"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Initialize config, JWT, and Firebase when the app starts
func init() {
	config.LoadEnv()
	service.InitJWTSecret()
	service.InitFirebaseClient()
}

func main() {
	// Load PORT from .env or default to 5000
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Set Gin mode dynamically based on .env
	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}

	// Initialize database connection
	database.Connect()
	config.InitRedis()
	config.InitRabbitMQ()
	workers.StartBatchFlusher()
	workers.StartLogBatchFlusher()

	// Start consumers
	go consumer.ConsumeQueue(config.RabbitMQConn, "logs", handlers.HandleLogEvent)
	go consumer.ConsumeQueue(config.RabbitMQConn, "location_updates", handlers.HandleLocationMessage)

	// Run migrations
	scripts.HandleScriptCommands()

	// Setup Gin router
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// Allow requests from the frontend
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Initialize repository

	// Register WebSocket routes
	websocket.RegisterRoutes(r, service.JwtSecret) // Make sure JwtSecret is exported in config

	// Register HTTP routes
	r.Static("/uploads", "./uploads")

	// temporary solution for the simulation
	vehicleRepo := &repository.VehicleRepository{}
	vehicleSvc := &service.VehicleService{Repo: vehicleRepo}
	vehicleCtrl := &controller.VehicleController{Service: vehicleSvc}
	PathRepo := &repository.PathRepository{}
	pathSvc := &service.PathService{Repo: PathRepo}
	pathCtrl := &controller.PathController{Service: pathSvc}

	r.GET("/simulation/vehicle", vehicleCtrl.GetAllVehicles) // Get all vehicles
	r.GET("/simulation/path", pathCtrl.GetAllPaths)          // Get all vehicles
	// temporary solution for the simulation

	routes.RegisterRoutes(r)

	// Start the server
	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Server is running on %s in %s mode", serverAddress, gin.Mode())
	if err := r.Run(serverAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
