package controller

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteController struct {
	Service *service.RouteService
}

// CreateRoute handles HTTP requests to create a new route
func (rc *RouteController) CreateRoute(c *gin.Context) {
	var route models.Route
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdRoute, err := rc.Service.CreateRoute(&route)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdRoute)
}

// UpdateRoute handles HTTP requests to update an existing route
func (rc *RouteController) UpdateRoute(c *gin.Context) {
	var route models.Route
	routeID := c.Param("id")

	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRoute, err := rc.Service.UpdateRoute(&route, utils.StringToUint(routeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRoute)
}

// DeleteRoute handles HTTP requests to delete a route
func (rc *RouteController) DeleteRoute(c *gin.Context) {
	routeID := c.Param("id")

	deletedRoute, err := rc.Service.DeleteRouteByID(utils.StringToUint(routeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedRoute)
}

// FindRouteByID handles HTTP requests to get a route by ID
func (rc *RouteController) FindRouteByID(c *gin.Context) {
	routeID := c.Param("id")

	route, err := rc.Service.FindRouteByID(utils.StringToUint(routeID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, route)
}

// GetAllRoutes handles HTTP requests to get all routes
func (rc *RouteController) GetAllRoutes(c *gin.Context) {
	routes, err := rc.Service.GetAllRoutes()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, routes)
}
