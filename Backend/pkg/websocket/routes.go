package websocket

import (
	"kassech/backend/pkg/repository"
	"kassech/backend/pkg/websocket/handlers"
	"kassech/backend/pkg/websocket/middleware"
	"kassech/backend/pkg/websocket/server"
	"kassech/backend/pkg/websocket/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, jwtSecret string) {
	// Initialize components

	userRepo := repository.UserRepository{} // Use concrete implementation
	passengerRepo := repository.PassengerRepository{}

	connManager := server.NewConnectionManager()
	auth := middleware.NewWebSocketAuth(jwtSecret)

	statusService := service.NewStatusService(userRepo)
	passengerService := service.NewPassengerService(passengerRepo)

	statusHandler := handlers.NewStatusHandler(connManager, statusService, auth)
	passengerHandler := handlers.NewPassengerHandler(connManager, passengerService, auth)

	// Register WebSocket routes
	router.GET("/ws/status", func(c *gin.Context) {
		statusHandler.HandleConnection(c.Writer, c.Request)
	})

	router.GET("/ws/passengers", func(c *gin.Context) {
		passengerHandler.HandleConnection(c.Writer, c.Request)
	})
}
