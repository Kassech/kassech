package controller

import (
	"errors"
	"fmt"
	"kassech/backend/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DriverDeligationController struct {
	Service *service.DriverDeligationService
}

type AssignDriverRequest struct {
	DriverID  uint `json:"driver_id" binding:"required"`
	VehicleID uint `json:"vehicle_id" binding:"required"`
}

func (c *DriverDeligationController) AssignDriver(ctx *gin.Context) {
	var req AssignDriverRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assignment, err := c.Service.AssignDriverToVehicle(req.DriverID, req.VehicleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, assignment)
}

func (c *DriverDeligationController) GetActiveDriver(ctx *gin.Context) {
	vehicleID := ctx.Param("vehicleId")
	var vehicleIDUint uint
	if _, err := fmt.Sscanf(vehicleID, "%d", &vehicleIDUint); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid vehicle ID"})
		return
	}

	assignment, err := c.Service.GetActiveDriver(vehicleIDUint)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "no active driver found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, assignment)
}
