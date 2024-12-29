package models

import "gorm.io/gorm"

// Route represents a connection between two stations
type Route struct {
	gorm.Model
	LocationA uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:LocationA;references:ID"`
	LocationB uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:LocationB;references:ID"`
}

