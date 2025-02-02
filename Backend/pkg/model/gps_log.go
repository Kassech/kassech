package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type VehicleGPSLog struct {
	gorm.Model
	VehicleID uint    `gorm:"not null" validate:"required" json:"vehicle_id"`
	Latitude  float64 `gorm:"not null" validate:"required,gt=0,lt=90" json:"lat"`
	Longitude float64 `gorm:"not null" validate:"required,gt=-180,lt=180" json:"lon"`
	PathID    uint    `validate:"omitempty" json:"path_id,omitempty"`
}

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
		// Handle string timestamps like "1709464702.123"
		ts, err := strconv.ParseFloat(t, 64)
		if err != nil {
			return fmt.Errorf("invalid timestamp format: %v", err)
		}
		sec := int64(ts)
		nsec := int64((ts - float64(sec)) * 1e9)
		v.CreatedAt = time.Unix(sec, nsec).UTC()
	case float64:
		// Handle numeric timestamps directly (e.g., 1709464702.123)
		sec := int64(t)
		nsec := int64((t - float64(sec)) * 1e9)
		v.CreatedAt = time.Unix(sec, nsec).UTC()
	default:
		return fmt.Errorf("unsupported timestamp type: %T", t)
	}

	return nil
}

func (v *VehicleGPSLog) Validate() error {
	return validator.New().Struct(v)
}
