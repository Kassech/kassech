package controller

import (
	"fmt"
	"kassech/backend/pkg/constants"
	"kassech/backend/pkg/domain"
	"kassech/backend/pkg/mapper"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	Service *service.VehicleService
}

// CreateVehicle handles HTTP requests to create a new vehicle
func (vc *VehicleController) CreateVehicle(c *gin.Context) {
	var vehicle domain.VehicleFormData
	if err := c.ShouldBind(&vehicle); err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("paths")
	paths := map[*multipart.FileHeader]string{
		vehicle.CarPicture: constants.CarPictureDirectory,
		vehicle.Bollo:      constants.CarBolloPath,
		vehicle.Insurance:  constants.CarInsurancePath,
		vehicle.Libre:      constants.CarLibrePath,
	}
	fmt.Println("paths")

	assignFuncs := map[*multipart.FileHeader]func(path string){
		vehicle.CarPicture: func(path string) { vehicle.CarPicturePath = &path },
		vehicle.Bollo:      func(path string) { vehicle.BolloPath = &path },
		vehicle.Insurance:  func(path string) { vehicle.InsurancePath = &path },
		vehicle.Libre:      func(path string) { vehicle.LibrePath = &path },
	}

	fmt.Println("paths")
	for file, dir := range paths {
		if file != nil {
			if err := utils.UploadAndAssignPath(file, dir, assignFuncs[file]); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save %s", dir)})
				return
			}
		}
	}
	fmt.Println("paths")
	vehicleModel, err := mapper.ToGormVehicle(&vehicle)
	fmt.Println("paths")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	createdVehicle, err := vc.Service.CreateVehicle(vehicleModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdVehicle)
}

// UpdateVehicle handles HTTP requests to update an existing vehicle
func (vc *VehicleController) UpdateVehicle(c *gin.Context) {
	var vehicle domain.VehicleFormData
	vehicleID := c.Param("id")

	if err := c.ShouldBind(&vehicle); err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paths := map[*multipart.FileHeader]string{
		vehicle.CarPicture: constants.CarPictureDirectory,
		vehicle.Bollo:      constants.CarBolloPath,
		vehicle.Insurance:  constants.CarInsurancePath,
		vehicle.Libre:      constants.CarLibrePath,
	}

	assignFuncs := map[*multipart.FileHeader]func(path string){
		vehicle.CarPicture: func(path string) { vehicle.CarPicturePath = &path },
		vehicle.Bollo:      func(path string) { vehicle.BolloPath = &path },
		vehicle.Insurance:  func(path string) { vehicle.InsurancePath = &path },
		vehicle.Libre:      func(path string) { vehicle.LibrePath = &path },
	}

	for file, dir := range paths {
		if file != nil {
			if err := utils.UploadAndAssignPath(file, dir, assignFuncs[file]); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save %s", dir)})
				return
			}
		}
	}

	vehicleModel, err := mapper.ToGormVehicle(&vehicle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	vehicleIDUint, err := utils.StringToUint(vehicleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	updatedVehicle, err := vc.Service.UpdateVehicle(vehicleModel, vehicleIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedVehicle)
}

// DeleteVehicle handles HTTP requests to delete a vehicle
func (vc *VehicleController) DeleteVehicle(c *gin.Context) {
	vehicleID := c.Param("id")
	vehicleIDUint, err := utils.StringToUint(vehicleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	forceDelete, err := utils.GetForceDeleteFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedVehicle, err := vc.Service.DeleteVehicleByID(vehicleIDUint, forceDelete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedVehicle)
}

// FindVehicleByID handles HTTP requests to get a vehicle by ID
func (vc *VehicleController) FindVehicleByID(c *gin.Context) {
	vehicleID := c.Param("id")
	vehicleIDUint, err := utils.StringToUint(vehicleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	vehicle, err := vc.Service.FindVehicleByID(vehicleIDUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vehicle)
}

// GetAllVehicles handles HTTP requests to get all vehicles
// TODO: fix the role based query
func (vc *VehicleController) GetAllVehicles(c *gin.Context) {
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

	search := c.Query("search")
	ownerID := c.Query("owner_id")
	typeID := c.Query("type_id")
	roles, ok := c.Get("role")
	if !ok {
		return
	}

	if utils.Contains(roles.([]string), constants.DriverRoleName) || utils.Contains(roles.([]string), constants.OwnerRoleName) {
		ownerID = fmt.Sprint(c.Get("userID"))
	}

	vehicles, total, err := vc.Service.GetAllVehicles(page, perPage, search, ownerID, typeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  vehicles,
		"total": total,
	})
}
