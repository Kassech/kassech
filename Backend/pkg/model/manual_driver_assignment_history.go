package models

import "gorm.io/gorm"

type ManualDriverAssignmentHistory struct {
	gorm.Model
	DriverID        uint   `gorm:"not null"` // Foreign key to Driver
	VehicleID       uint   `gorm:"not null"` // Foreign key to Vehicle
	AssignedByID    uint   `gorm:"not null"` // Foreign key to User (assigned by)
	AssignmentDate  string `gorm:"size:20;not null"`
	AssignmentNotes string `gorm:"type:text"` // Optional notes field
}
