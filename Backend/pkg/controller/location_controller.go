package controller

import (
	"fmt"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
	Service *service.LocationService
}

// CreateLocation handles HTTP requests to create a new location
func (lc *LocationController) CreateLocation(c *gin.Context) {
	var location models.VehicleGPSLog
	if err := c.ShouldBind(&location); err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := location.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdLocation, err := lc.Service.CreateLocation(&location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdLocation)
}

// DeleteLocation handles HTTP requests to delete a location
func (lc *LocationController) DeleteLocation(c *gin.Context) {
	locationID := c.Param("id")
	locationIDUint, err := utils.StringToUint(locationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	forceDelete, err := utils.GetForceDeleteFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedLocation, err := lc.Service.DeleteLocationByID(locationIDUint, forceDelete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedLocation)
}

// FindLocationByID handles HTTP requests to get a location by ID
func (lc *LocationController) FindLocationByID(c *gin.Context) {
	locationID := c.Param("id")
	locationIDUint, err := utils.StringToUint(locationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	location, err := lc.Service.FindLocationByID(locationIDUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, location)
}

// GetAllLocations handles HTTP requests to get all locations
// TODO: fix the role based query
func (lc *LocationController) GetAllLocations(c *gin.Context) {
	page, err := utils.GetPageFromQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	perPage, err := utils.GetPerPageFromQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	vehicleIDStr := c.Query("vehicle_id")
	pathIDStr := c.Query("path_id")
	search := c.Query("search")

	vehicleIDUint, err := utils.StringToUint(vehicleIDStr)
	if err != nil {
		vehicleIDUint = 0
	}
	pathIDUint, err := utils.StringToUint(pathIDStr)
	if err != nil {
		pathIDUint = 0
	}
	locations, total, err := lc.Service.GetAllLocations(page, perPage, search, vehicleIDUint, pathIDUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  locations,
		"total": total,
	})
}
