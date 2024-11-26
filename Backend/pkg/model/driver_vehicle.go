package models

import "gorm.io/gorm"

type DriverVehicle struct {
	gorm.Model
	DriverID  uint   `gorm:"not null"` // Foreign key to Driver
	VehicleID uint   `gorm:"not null"` // Foreign key to Vehicle
	Status    string `gorm:"size:20;not null;default:'active'"`
}
