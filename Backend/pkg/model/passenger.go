package models

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model
	PathID uint   `gorm:"not null"`
	Status string `gorm:"not null"`
}
