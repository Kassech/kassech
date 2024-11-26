package models

import "gorm.io/gorm"

type Path struct {
    gorm.Model
    RouteID       uint    `gorm:"not null"`
    PathName      string  `gorm:"not null"`
    DistanceKM    float64 `gorm:"not null"`
    EstimatedTime string  `gorm:"not null"`
    IsActive      bool    `gorm:"default:true"`
}
