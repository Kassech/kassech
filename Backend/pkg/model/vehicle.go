package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	Type          uint   `gorm:"not null"`
	LicenseNumber string `gorm:"size:20;not null"`
	VIN           string `gorm:"size:20;unique;not null"`
	Make          string `gorm:"not null"`
	Year          uint   `gorm:"not null"`
	Color         string
	OwnerID       uint   `gorm:"not null"`
	OwnerSince    string `gorm:"not null"`
}
