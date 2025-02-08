package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"

	"gorm.io/gorm"
)

type DriverDeligationRepository struct{}

func (r *DriverDeligationRepository) Create(driverVehicle *models.DriverVehicle) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// Deactivate all active drivers for this vehicle
		if err := tx.Model(&models.DriverVehicle{}).
			Where("vehicle_id = ? AND status = ?", driverVehicle.VehicleID, "active").
			Update("status", "inactive").Error; err != nil {
			return err
		}

		// Create new active assignment
		return tx.Create(driverVehicle).Error
	})
}

func (r *DriverDeligationRepository) GetActiveDriver(vehicleID uint) (*models.DriverVehicle, error) {
	var assignment models.DriverVehicle
	err := database.DB.Where("vehicle_id = ? AND status = ?", vehicleID, "active").First(&assignment).Error
	return &assignment, err
}
