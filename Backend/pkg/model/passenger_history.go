package models

import "gorm.io/gorm"

type PassengerHistory struct {
	gorm.Model
	PassengerID uint   `gorm:"not null"` // Foreign key to Passenger
	VehicleID   uint   `gorm:"not null"` // Foreign key to Vehicle
	RouteID     uint   `gorm:"not null"` // Foreign key to Route
	TravelDate  string `gorm:"size:20;not null"`
}
