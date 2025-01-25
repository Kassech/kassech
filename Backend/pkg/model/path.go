package models

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Path struct {
	gorm.Model
	RouteID       uint          `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"route_id"`
	Route         Route         `gorm:"foreignKey:RouteID" json:"route"`
	PathName      string        `gorm:"not null" validate:"required" json:"path_name"`
	DistanceKM    float64       `gorm:"not null" validate:"required,gt=0" json:"distance_km"`
	EstimatedTime time.Duration `gorm:"not null" validate:"required,gt=0" json:"estimated_time"` // Using time.Duration
	IsActive      bool          `gorm:"default:true" json:"is_active"`
}

// MarshalJSON to convert time.Duration to string for JSON response
func (p *Path) MarshalJSON() ([]byte, error) {
	type Alias Path
	return json.Marshal(&struct {
		EstimatedTime string `json:"estimated_time"`
		*Alias
	}{
		EstimatedTime: p.EstimatedTime.String(), // Convert time.Duration to string
		Alias:         (*Alias)(p),
	})
}

// UnmarshalJSON to convert string back to time.Duration for JSON request
func (p *Path) UnmarshalJSON(data []byte) error {
	type Alias Path
	aux := &struct {
		EstimatedTime string `json:"estimated_time"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Convert string to time.Duration
	duration, err := time.ParseDuration(aux.EstimatedTime)
	if err != nil {
		return err
	}
	p.EstimatedTime = duration
	return nil
}

// BeforeCreate hook to validate the struct before saving to the database
func (p *Path) BeforeCreate(tx *gorm.DB) error {
	v := validator.New()
	if err := v.Struct(p); err != nil {
		return err
	}
	return nil
}
