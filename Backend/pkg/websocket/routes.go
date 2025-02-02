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
	locationRepo := repository.LocationRepository{}

	connManager := server.NewConnectionManager()
	auth := middleware.NewWebSocketAuth(jwtSecret)

	statusService := service.NewStatusService(userRepo)
	passengerService := service.NewPassengerService(passengerRepo)
	locationService := service.NewLocationService(locationRepo)

	statusHandler := handlers.NewStatusHandler(connManager, statusService, auth)
	passengerHandler := handlers.NewPassengerHandler(connManager, passengerService, auth)
	locationHandler := handlers.NewLocationHandler(connManager, locationService, auth)

	// Register WebSocket routes
	router.GET("/ws/status", func(c *gin.Context) {
		statusHandler.HandleConnection(c.Writer, c.Request)
	})
	router.GET("/ws/location", func(c *gin.Context) {
		locationHandler.HandleConnection(c.Writer, c.Request)
	})

	router.GET("/ws/passengers", func(c *gin.Context) {
		passengerHandler.HandleConnection(c.Writer, c.Request)
	})
}
