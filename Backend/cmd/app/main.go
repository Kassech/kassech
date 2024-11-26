package main

import (
	"kassech/backend/pkg/config"
	"kassech/backend/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load environment variables
	config.LoadEnv()

}
func main() {

	// Initialize DB connection
	db, err := database.Connect()
	database.Migrate(db)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	// Initialize Gin router
	r := gin.Default()
	r.Run(":8080")

}
