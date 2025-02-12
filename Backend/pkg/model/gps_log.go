package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// VehicleGPSLog stores GPS coordinates using PostgreSQL's PostGIS extension
type VehicleGPSLog struct {
	gorm.Model
	VehicleID uint      `gorm:"not null" json:"vehicle_id"`
	Latitude  float64   `gorm:"-" json:"lat"`
	Longitude float64   `gorm:"-" json:"lon"`
	Location  string    `gorm:"type:geometry(Point,4326);not null"`
	PathID    uint      `json:"path_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
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

// AfterFind hook to extract lat/lon from Location
func (v *VehicleGPSLog) AfterFind(tx *gorm.DB) error {
	if v.Location != "" {
		var lon, lat float64
		// Log the location value for debugging
		fmt.Printf("Raw Location: %s\n", v.Location)

		// Ensure the format is "POINT(lon lat)" without any extra characters
		trimmed := strings.TrimSpace(v.Location)
		if strings.HasPrefix(trimmed, "POINT(") && strings.HasSuffix(trimmed, ")") {
			trimmed = trimmed[6 : len(trimmed)-1] // Remove "POINT(" and ")"
			coords := strings.Split(trimmed, " ")
			if len(coords) == 2 {
				// Convert to float values
				var err error
				lon, err = strconv.ParseFloat(coords[0], 64)
				if err != nil {
					return fmt.Errorf("failed to parse longitude: %v", err)
				}
				lat, err = strconv.ParseFloat(coords[1], 64)
				if err != nil {
					return fmt.Errorf("failed to parse latitude: %v", err)
				}

				v.Longitude = lon
				v.Latitude = lat
				return nil
			}
		}

		return fmt.Errorf("failed to parse Location: invalid format")
	}
	return nil
}

// Validate function
func (v *VehicleGPSLog) Validate() error {
	return validator.New().Struct(v)
}
