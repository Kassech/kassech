package models

import "gorm.io/gorm"

// Route represents a connection between two stations
type Route struct {
	gorm.Model
	LocationA uint    `gorm:"not null"`
	StationA  Station `gorm:"foreignKey:LocationA;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LocationB uint    `gorm:"not null"`
	StationB  Station `gorm:"foreignKey:LocationB;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
