package models

import "gorm.io/gorm"

type DriverAssignmentHistory struct {
	gorm.Model
	DriverID        uint   `gorm:"not null"`     // Foreign key to Driver
	VehicleID       uint   `gorm:"not null"`     // Foreign key to Vehicle
	AssignedByID    *uint  `gorm:"default:null"` // Nullable, only set for manual assignments
	AssignmentDate  string `gorm:"size:20;not null"`
	AssignmentNotes string `gorm:"type:text"`        // Optional notes field for manual assignments
	AlgorithmUsed   string `gorm:"size:255"`         // Optional, only set for auto assignments
	AssignmentType  string `gorm:"size:10;not null"` // "manual" or "auto"
}
