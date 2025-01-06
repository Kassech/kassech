package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"kassech/backend/pkg/config"
	"kassech/backend/pkg/database"
	routes "kassech/backend/pkg/delivery/http"
	"kassech/backend/pkg/delivery/socket"
	"kassech/backend/pkg/service"
	scripts "kassech/backend/script"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() { // Load environment variables
	config.LoadEnv()
	service.InitJWTSecret()
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
	// Run migrations
	scripts.HandleScriptCommands()

	// Setup Gin router
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Register routes
	socket.RegisterRoutes(r)
	r.Static("/uploads", "./uploads")
	routes.RegisterRoutes(r)
	// r.Use(middleware.AuthMiddleware())

	// Start the server
	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Server is running on %s in %s mode", serverAddress, gin.Mode())
	if err := r.Run(serverAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
