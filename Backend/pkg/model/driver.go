package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	UserID        uint
	DriverLicense string `gorm:"not null"`
	Status        string
}
