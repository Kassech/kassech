package models

import "gorm.io/gorm"

type Route struct {
    gorm.Model
    LocationA uint `gorm:"not null"`
    LocationB uint `gorm:"not null"`
}
