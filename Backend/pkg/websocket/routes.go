package websocket

import (
	"kassech/backend/pkg/config"
	"kassech/backend/pkg/repository"
	"kassech/backend/pkg/websocket/handlers"
	"kassech/backend/pkg/websocket/middleware"
	"kassech/backend/pkg/websocket/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, jwtSecret string) {
	// Initialize components

	userRepo := repository.UserRepository{} // Use concrete implementation
	passengerRepo := repository.PassengerRepository{}
	locationRepo := repository.LocationRepository{}
	destinationRepo := repository.DestinationRepository{}

	auth := middleware.NewWebSocketAuth(jwtSecret)

	statusService := service.NewStatusService(userRepo)
	passengerService := service.NewPassengerService(passengerRepo)
	locationService := service.NewLocationService(locationRepo)
	destinationService := service.NewDestinationService(destinationRepo)

	statusHandler := handlers.NewStatusHandler(config.ConnManager, statusService, auth)
	passengerHandler := handlers.NewPassengerHandler(config.ConnManager, passengerService, auth)
	locationHandler := handlers.NewLocationHandler(config.ConnManager, locationService, auth)
	destinationHandler := handlers.NewDestinationHandler(config.ConnManager, destinationService, auth)
	// Register WebSocket routes
	router.GET("/ws/status", func(c *gin.Context) {
		statusHandler.HandleConnection(c.Writer, c.Request)
	})
	router.GET("/ws/location", func(c *gin.Context) {
		locationHandler.HandleConnection(c.Writer, c.Request)
	})
	router.GET("/ws/destination", func(c *gin.Context) {
		destinationHandler.HandleConnection(c.Writer, c.Request)
	})
	router.GET("/ws/passengers", func(c *gin.Context) {
		passengerHandler.HandleConnection(c.Writer, c.Request)
	})
}
