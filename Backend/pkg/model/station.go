package models

import "gorm.io/gorm"

type Station struct {
    gorm.Model
    LocationName string  `gorm:"not null"`
    Latitude     float64 `gorm:"not null"`
    Longitude    float64 `gorm:"not null"`
}
