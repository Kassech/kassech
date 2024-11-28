package main

import (
	"fmt"
	"log"
	"os"

	"kassech/backend/pkg/config"
	"kassech/backend/pkg/database"
	"kassech/backend/pkg/delivery/http"
	"kassech/backend/pkg/service"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load environment variables
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

	// database.Migrate()
	// database.SeedDB()

	// Setup Gin router
	r := gin.Default()

	// Register routes
	http.RegisterRoutes(r)

	// Start the server
	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Server is running on %s in %s mode", serverAddress, gin.Mode())
	if err := r.Run(serverAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
