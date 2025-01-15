package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type DriverRepository struct{}

// Create inserts a new driver
func (dr *DriverRepository) Create(driver *models.Driver) (*models.Driver, error) {
	err := database.DB.Create(driver).Error
	if err != nil {
		return nil, err
	}
	return driver, nil
}

// FindByEmailOrPhone searches for a driver by either email or phone number
func (dr *DriverRepository) FindByEmailOrPhone(email string, phone string) (*models.Driver, error) {
	var driver models.Driver
	err := database.DB.Where("email = ? OR phone_number = ?", email, phone).First(&driver).Error
	if err != nil {
		return nil, err
	}
	return &driver, nil
}

// FindByID fetches a driver by their unique ID
func (dr *DriverRepository) FindByID(driverID uint) (*models.Driver, error) {
	var driver models.Driver
	err := database.DB.First(&driver, driverID).Error
	if err != nil {
		return nil, err
	}
	return &driver, nil
}
