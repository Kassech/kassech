package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// VehicleGPSLog stores GPS coordinates using PostgreSQL's PostGIS extension
type VehicleGPSLog struct {
	gorm.Model
	VehicleID uint    `gorm:"not null" validate:"required" json:"vehicle_id"`
	Latitude  float64 `gorm:"-" json:"lat"`                       // Ignored in DB, used for JSON
	Longitude float64 `gorm:"-" json:"lon"`                       // Ignored in DB, used for JSON
	Location  string  `gorm:"type:geometry(Point,4326);not null"` // Stores lat/lon as PG geo data
	PathID    uint    `validate:"omitempty" json:"path_id,omitempty"`
}

// BeforeCreate hook to store lat/lon as PostGIS POINT
func (v *VehicleGPSLog) BeforeCreate(tx *gorm.DB) error {
	v.Location = fmt.Sprintf("SRID=4326;POINT(%f %f)", v.Longitude, v.Latitude)
	return nil
}

// UnmarshalJSON to handle timestamps
func (v *VehicleGPSLog) UnmarshalJSON(data []byte) error {
	type Alias VehicleGPSLog
	aux := &struct {
		CreatedAt interface{} `json:"created_at"`
		*Alias
	}{
		Alias: (*Alias)(v),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch t := aux.CreatedAt.(type) {
	case string:
		ts, err := strconv.ParseFloat(t, 64)
		if err != nil {
			return fmt.Errorf("invalid timestamp format: %v", err)
		}
		sec := int64(ts)
		nsec := int64((ts - float64(sec)) * 1e9)
		v.CreatedAt = time.Unix(sec, nsec).UTC()
	case float64:
		sec := int64(t)
		nsec := int64((t - float64(sec)) * 1e9)
		v.CreatedAt = time.Unix(sec, nsec).UTC()
	default:
		return fmt.Errorf("unsupported timestamp type: %T", t)
	}

	return nil
}

// Validate function
func (v *VehicleGPSLog) Validate() error {
	return validator.New().Struct(v)
}
