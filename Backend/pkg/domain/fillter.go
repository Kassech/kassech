package domain

import "time"

type GPSLogFilter struct {
	VehicleIDs []uint    `json:"vehicle_ids" form:"vehicle_id"`
	PathIDs    []uint    `json:"path_ids" form:"path_id"`
	StartTime  time.Time `json:"start_time" form:"start_time"`
	EndTime    time.Time `json:"end_time" form:"end_time"`
	CenterLat  float64   `json:"lat" form:"lat"`
	CenterLon  float64   `json:"lon" form:"lon"`
	Radius     float64   `json:"radius" form:"radius"` // in meters
	Page       int       `json:"page" form:"page"`
	PerPage    int       `json:"per_page" form:"per_page"`
}
