package models

import "gorm.io/gorm"

type VehicleType struct {
	gorm.Model
	TypeName    string `gorm:"size:255;not null;unique"`
	Capacity    uint   `gorm:"not null"`
	Description string `gorm:"type:text"` // Optional description of the vehicle type
}
