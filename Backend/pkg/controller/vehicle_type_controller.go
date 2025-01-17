package controller

import (
	"fmt"
	"kassech/backend/pkg/domain"
	"kassech/backend/pkg/mapper"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehicleTypeController struct {
	Service *service.VehicleTypeService
}

// CreateVehicleType handles HTTP requests to create a new vehicle type
func (vtc *VehicleTypeController) CreateVehicleType(c *gin.Context) {
	var vehicleType domain.VehicleTypeFormData
	if err := c.ShouldBind(&vehicleType); err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicleTypeModel, err := mapper.ToGormVehicleType(&vehicleType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	createdVehicleType, err := vtc.Service.CreateVehicleType(vehicleTypeModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdVehicleType)
}

// UpdateVehicleType handles HTTP requests to update an existing vehicle type
func (vtc *VehicleTypeController) UpdateVehicleType(c *gin.Context) {
	var vehicleType models.VehicleType
	vehicleTypeID := c.Param("id")

	if err := c.ShouldBindJSON(&vehicleType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vehicleTypeIDUint, err := utils.StringToUint(vehicleTypeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle type ID"})
		return
	}
	updatedVehicleType, err := vtc.Service.UpdateVehicleType(&vehicleType, vehicleTypeIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedVehicleType)
}

// DeleteVehicleType handles HTTP requests to delete a vehicle type
func (vtc *VehicleTypeController) DeleteVehicleType(c *gin.Context) {
	vehicleTypeID := c.Param("id")
	vehicleTypeIDUint, err := utils.StringToUint(vehicleTypeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle type ID"})
		return
	}

	forceDelete, err := utils.GetForceDeleteFromHeader(c)
	fmt.Println("forceDelete:", forceDelete)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedVehicleType, err := vtc.Service.DeleteVehicleTypeByID(vehicleTypeIDUint, forceDelete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedVehicleType)
}

// FindVehicleTypeByID handles HTTP requests to get a vehicle type by ID
func (vtc *VehicleTypeController) FindVehicleTypeByID(c *gin.Context) {
	vehicleTypeID := c.Param("id")
	vehicleTypeIDUint, err := utils.StringToUint(vehicleTypeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle type ID"})
		return
	}
	forceDelete, _ := utils.GetForceDeleteFromHeader(c)
	fmt.Println("forceDelete:", forceDelete)

	vehicleType, err := vtc.Service.FindVehicleTypeByID(vehicleTypeIDUint, forceDelete)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vehicleType)
}

// GetAllVehicleTypes handles HTTP requests to get all vehicle types
func (vtc *VehicleTypeController) GetAllVehicleTypes(c *gin.Context) {
	vehicleTypes, err := vtc.Service.GetAllVehicleTypes()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vehicleTypes)
}
