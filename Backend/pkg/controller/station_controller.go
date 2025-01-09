package controller

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StationController struct {
	Service *service.StationService
}

// CreateStation handles HTTP requests to create a new role
func (rc *StationController) CreateStation(c *gin.Context) {
	var role models.Station
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdStation, err := rc.Service.CreateStation(&role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdStation)
}

// UpdateStation handles HTTP requests to update an existing role
func (rc *StationController) UpdateStation(c *gin.Context) {
	var role models.Station
	roleID := c.Param("id")

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roleIdUint, err := utils.StringToUint(roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	updatedStation, err := rc.Service.UpdateStation(&role, roleIdUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedStation)
}

// DeleteStation handles HTTP requests to delete a role
func (rc *StationController) DeleteStation(c *gin.Context) {
	roleID := c.Param("id")
	roleIdUint, err := utils.StringToUint(roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	deletedStation, err := rc.Service.DeleteStationByID(roleIdUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedStation)
}

// FindStationByID handles HTTP requests to get a role by ID
func (rc *StationController) FindStationByID(c *gin.Context) {
	roleID := c.Param("id")
	userIdUint, err := utils.StringToUint(roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	role, err := rc.Service.FindStationByID(userIdUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// GetAllStation handles HTTP requests to get a role by ID
func (rc *StationController) GetAllStations(c *gin.Context) {

	role, err := rc.Service.GetAllStations()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}
