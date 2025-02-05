package models

import "gorm.io/gorm"

// Route represents a connection between two stations
type Route struct {
	gorm.Model
	LocationA uint    `gorm:"not null" json:"location_a"`
	StationA  Station `gorm:"foreignKey:LocationA;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"station_a"`
	LocationB uint    `gorm:"not null" json:"location_b"`
	StationB  Station `gorm:"foreignKey:LocationB;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"station_b"`
}
