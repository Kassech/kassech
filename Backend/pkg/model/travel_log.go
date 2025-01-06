package models

import "gorm.io/gorm"

type TravelLog struct {
	gorm.Model
	DriverID   uint    `gorm:"not null"`
	CarID      uint    `gorm:"not null"`
	PathID     uint    `gorm:"not null"`
	StartTime  string  `gorm:"not null"`
	EndTime    string  `gorm:"not null"`
	DistanceKM float64 `gorm:"not null"`
}
