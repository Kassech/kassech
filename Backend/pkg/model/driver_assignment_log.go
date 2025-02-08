package models

import (
	"time"

	"gorm.io/gorm"
)

type DriverAssignmentLog struct {
	gorm.Model
	VehicleID      uint      `gorm:"not null"`
	DriverID       uint      `gorm:"not null"`
	AssignedAt     time.Time `gorm:"not null"`
	Action         string    `gorm:"size:20;not null"`
	PerformedBy    string    `gorm:"size:100;not null"`
	PreviousDriver *uint
}
