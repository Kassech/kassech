package models

import "gorm.io/gorm"

type AutoDriverAssignmentHistory struct {
	gorm.Model
	DriverID       uint   `gorm:"not null"` // Foreign key to Driver
	VehicleID      uint   `gorm:"not null"` // Foreign key to Vehicle
	AssignedDate   string `gorm:"size:20;not null"`
	AlgorithmUsed  string `gorm:"size:255"` // Information about the assignment algorithm
}
