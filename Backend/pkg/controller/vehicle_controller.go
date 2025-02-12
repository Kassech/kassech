package controller

import (
	"fmt"
	"kassech/backend/pkg/constants"
	"kassech/backend/pkg/domain"
	"kassech/backend/pkg/mapper"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

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
		log.Printf("Error getting page: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	perPage, err := utils.GetPerPageFromQuery(c)
	if err != nil {
		log.Printf("Error getting per page: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if err != nil || limit < 1 {
		log.Printf("Invalid limit, defaulting to 10: %v", err)
		limit = 100
	}

	search := c.Query("search")
	ownerID := c.Query("owner_id")
	typeID := c.Query("type_id")
	// roles, ok := c.Get("role")
	// if !ok {
	// 	log.Println("Roles not found in context")
	// 	return
	// }

	// if utils.Contains(roles.([]string), constants.DriverRoleName) || utils.Contains(roles.([]string), constants.OwnerRoleName) {
	// 	ownerID = fmt.Sprint(c.Get("userID"))
	// }

	log.Printf("Fetching vehicles with search: %s, ownerID: %s, typeID: %s", search, ownerID, typeID)
	vehicles, total, err := vc.Service.GetAllVehicles(page, perPage, search, ownerID, typeID)
	if err != nil {
		log.Printf("Error fetching vehicles: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  vehicles,
		"total": total,
	})
}

// UpdateVehicleStatus updates the status of a vehicle
func (vc *VehicleController) UpdateVehicleStatus(c *gin.Context) {
	vehicleID := c.Param("id")
	vehicleIDUint, err := utils.StringToUint(vehicleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle ID"})
		return
	}
	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}
	if err := vc.Service.UpdateVehicleStatus(vehicleIDUint, status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vehicle status updated successfully"})
}

func (vc *VehicleController) FilterGPSLogs(c *gin.Context) {
	var filter domain.GPSLogFilter

	// Parse vehicle IDs
	if vehicleIDs := c.QueryArray("vehicle_id"); len(vehicleIDs) > 0 {
		for _, idStr := range vehicleIDs {
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle ID"})
				return
			}
			filter.VehicleIDs = append(filter.VehicleIDs, uint(id))
		}
	}

	// Parse path IDs
	if pathIDs := c.QueryArray("path_id"); len(pathIDs) > 0 {
		for _, idStr := range pathIDs {
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid path ID"})
				return
			}
			filter.PathIDs = append(filter.PathIDs, uint(id))
		}
	}

	// Parse time range
	if startTime := c.Query("start_time"); startTime != "" {
		st, err := time.Parse(time.RFC3339, startTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
			return
		}
		filter.StartTime = st
	}
	if endTime := c.Query("end_time"); endTime != "" {
		et, err := time.Parse(time.RFC3339, endTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
			return
		}
		filter.EndTime = et
	}

	// Parse location filter
	if lat := c.Query("lat"); lat != "" {
		if lat, err := strconv.ParseFloat(lat, 64); err == nil {
			filter.CenterLat = lat
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
			return
		}
	}
	if lon := c.Query("lon"); lon != "" {
		if lon, err := strconv.ParseFloat(lon, 64); err == nil {
			filter.CenterLon = lon
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
			return
		}
	}
	if radius := c.Query("radius"); radius != "" {
		if radius, err := strconv.ParseFloat(radius, 64); err == nil {
			filter.Radius = radius
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid radius"})
			return
		}
	}

	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	filter.Page = page
	filter.PerPage = perPage

	logs, total, err := vc.Service.FilterGPSLogs(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  logs,
		"total": total,
	})
}
