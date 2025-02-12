package repository

import (
	"errors"
	"fmt"
	"kassech/backend/pkg/database"
	"kassech/backend/pkg/domain"
	models "kassech/backend/pkg/model"
	"time"

	"math"
	"sort"

	"gorm.io/gorm"
)

type VehicleRepository struct{}

// Create a new vehicle in the database
func (rr *VehicleRepository) Create(vehicle *models.Vehicle) (*models.Vehicle, error) {
	if err := database.DB.Create(vehicle).Error; err != nil {
		return nil, err
	}
	return vehicle, nil
}

// Find a vehicle by ID with active driver
func (rr *VehicleRepository) FindByID(vehicleID uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := database.DB.
		Preload("Owner").
		Preload("Type").
		Preload("Driver.User").
		First(&vehicle, vehicleID).Error

	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// GetAll vehicles with optional filters and active drivers
func (rr *VehicleRepository) GetAll(page, perPage int, search, ownerID, typeID string) ([]models.Vehicle, int64, error) {
	var vehicles []models.Vehicle
	var total int64

	query := database.DB.
		Preload("Owner").
		Preload("Type").
		Preload("Driver.User").
		Model(&models.Vehicle{})

	if search != "" {
		query = query.Where("license_number ILIKE ? OR vin ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if ownerID != "" {
		query = query.Where("owner_id = ?", ownerID)
	}

	if typeID != "" {
		query = query.Where("type_id = ?", typeID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset((page - 1) * perPage).Limit(perPage).Find(&vehicles).Error
	return vehicles, total, err
}

// Update an existing vehicle by ID
func (rr *VehicleRepository) Update(vehicle *models.Vehicle, vehicleID uint) (*models.Vehicle, error) {
	existingVehicle, err := rr.FindByID(vehicleID)
	if err != nil {
		return nil, err
	}

	if err := database.DB.Model(existingVehicle).Updates(vehicle).Error; err != nil {
		return nil, err
	}
	return existingVehicle, nil
}

// Delete a vehicle by ID
func (rr *VehicleRepository) DeleteByID(vehicleID uint, forceDelete bool) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	query := database.DB
	if forceDelete {
		query = query.Unscoped()
	}

	if err := query.First(&vehicle, vehicleID).Error; err != nil {
		return nil, err
	}

	if err := query.Delete(&vehicle).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}
func (vr *VehicleRepository) GetNearestAvailableCars(stationID uint, timeRangeInMinutes int) ([]models.Vehicle, error) {
	// Step 1: Get the station's coordinates
	var station models.Station
	if err := database.DB.First(&station, stationID).Error; err != nil {
		return nil, err
	}

	// Step 2: Get the latest GPS logs for all vehicles within the time range
	timeRange := time.Now().Add(-time.Duration(timeRangeInMinutes) * time.Minute)
	var gpsLogs []models.VehicleGPSLog
	subQuery := database.DB.
		Model(&models.VehicleGPSLog{}).
		Select("vehicle_id, MAX(created_at) as latest").
		Where("created_at > ?", timeRange).
		Group("vehicle_id")

	if err := database.DB.
		Joins("JOIN (?) AS latest_logs ON vehicle_gps_logs.vehicle_id = latest_logs.vehicle_id AND vehicle_gps_logs.created_at = latest_logs.latest", subQuery).
		Find(&gpsLogs).Error; err != nil {
		return nil, err
	}

	// Step 3: Fetch vehicles with preloaded Type and calculate distances
	type CarDistance struct {
		Vehicle  models.Vehicle
		Distance float64
	}

	var carsWithDistance []CarDistance
	for _, log := range gpsLogs {
		var vehicle models.Vehicle
		// Preload Type to ensure capacity is available
		if err := database.DB.
			Preload("Owner").
			Preload("Type").
			Preload("Driver.User").
			Where("status IN (?)", []string{"un-assigned", "active"}).
			First(&vehicle, log.VehicleID).Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Vehicle exists but has invalid status (or soft-deleted)
				fmt.Printf("Vehicle %d skipped: not found or invalid status", log.VehicleID)
			} else {
				// Other database errors
				fmt.Printf("Error fetching vehicle %d: %v", log.VehicleID, err)
			}
			continue
		}
		// Skip vehicles with invalid capacity
		if vehicle.Type.Capacity <= 0 {
			continue
		}

		distance := Haversine(station.Latitude, station.Longitude, log.Latitude, log.Longitude)
		carsWithDistance = append(carsWithDistance, CarDistance{
			Vehicle:  vehicle,
			Distance: distance,
		})
	}

	// Step 4: Sort by distance (ascending)
	sort.Slice(carsWithDistance, func(i, j int) bool {
		return carsWithDistance[i].Distance < carsWithDistance[j].Distance
	})

	// Extract sorted vehicles
	var sortedVehicles []models.Vehicle
	for _, cd := range carsWithDistance {
		sortedVehicles = append(sortedVehicles, cd.Vehicle)
	}

	return sortedVehicles, nil
}

// Haversine calculates the great-circle distance between two points
// on the Earth (specified in decimal degrees) using the Haversine formula.
// Returns distance in kilometers.
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // Earth's mean radius in kilometers

	// Convert degrees to radians
	degToRad := func(deg float64) float64 { return deg * math.Pi / 180 }

	lat1Rad := degToRad(lat1)
	lon1Rad := degToRad(lon1)
	lat2Rad := degToRad(lat2)
	lon2Rad := degToRad(lon2)

	// Differences in coordinates
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// Haversine formula components
	a := math.Pow(math.Sin(dLat/2), 2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Pow(math.Sin(dLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate the distance
	distance := earthRadius * c

	return distance
}

// GetAvailableCars retrieves available cars for a specified path.
func (vr *VehicleRepository) GetAvailableCars(pathID uint) ([]models.Vehicle, error) {
	var availableCars []models.Vehicle

	if err := database.DB.
		Where("path_id = ? AND is_available = ?", pathID, true).
		Find(&availableCars).Error; err != nil {
		return nil, err
	}

	return availableCars, nil
}

// In repository/vehicle_repository.go
func (vr *VehicleRepository) UpdateVehicleStatus(vehicleID uint, status string) error {
	var vehicle models.Vehicle

	// Fetch the vehicle
	if err := database.DB.First(&vehicle, vehicleID).Error; err != nil {
		return fmt.Errorf("vehicle not found: %v", err)
	}

	// Use the model's SetStatus method (with validation)
	if err := vehicle.SetStatus(status); err != nil {
		return err // Return validation error (e.g., invalid status)
	}

	// Save the updated status to the database
	if err := database.DB.Save(&vehicle).Error; err != nil {
		return fmt.Errorf("failed to update status: %v", err)
	}

	return nil
}

func GetLatestNearbyLocation(db *gorm.DB, currentLat, currentLon float64, maxDistance float64) (*models.VehicleGPSLog, error) {
	var latestLog models.VehicleGPSLog

	// Calculate the time 5 minutes ago
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)

	// Query to find the latest nearby location within the last 5 minutes
	err := db.Where("created_at >= ?", fiveMinutesAgo).
		Where("ST_DWithin(location::geography, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)", currentLon, currentLat, maxDistance).
		Order("created_at DESC").
		First(&latestLog).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no nearby location found within the last 5 minutes")
		}
		return nil, err
	}

	return &latestLog, nil
}

func (vr *VehicleRepository) FilterGPSLogs(filter domain.GPSLogFilter) ([]models.VehicleGPSLog, int64, error) {
	var logs []models.VehicleGPSLog
	var total int64

	query := database.DB.Model(&models.VehicleGPSLog{})

	// Apply filters
	if len(filter.VehicleIDs) > 0 {
		query = query.Where("vehicle_id IN (?)", filter.VehicleIDs)
	}
	if len(filter.PathIDs) > 0 {
		query = query.Where("path_id IN (?)", filter.PathIDs)
	}
	if !filter.StartTime.IsZero() {
		query = query.Where("created_at >= ?", filter.StartTime)
	}
	if !filter.EndTime.IsZero() {
		query = query.Where("created_at <= ?", filter.EndTime)
	}
	if filter.Radius > 0 && filter.CenterLat != 0 && filter.CenterLon != 0 {
		query = query.Where("ST_DWithin(location::geography, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)",
			filter.CenterLon, filter.CenterLat, filter.Radius)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	if filter.Page > 0 && filter.PerPage > 0 {
		offset := (filter.Page - 1) * filter.PerPage
		query = query.Offset(offset).Limit(filter.PerPage)
	}

	// Execute query
	if err := query.Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
