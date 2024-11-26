package models

import "gorm.io/gorm"

type VehicleGPSLog struct {
	gorm.Model
	VehicleID uint    `gorm:"not null"`
	Latitude  float64 `gorm:"not null"`
	Longitude float64 `gorm:"not null"`
	PathID    uint
}
