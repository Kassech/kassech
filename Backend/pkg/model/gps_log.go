package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type VehicleGPSLog struct {
	gorm.Model
	VehicleID uint    `gorm:"not null" validate:"required"`
	Latitude  float64 `gorm:"not null" validate:"required,gt=0,lt=90"`
	Longitude float64 `gorm:"not null" validate:"required,gt=-180,lt=180"`
	PathID    uint    `validate:"omitempty"`
}

// Validate will use the validator package to validate the VehicleGPSLog struct
func (v *VehicleGPSLog) Validate() error {
	return validator.New().Struct(v)
}

