package models

import "gorm.io/gorm"

type QueueManagerRoute struct {
	gorm.Model
	QueueManagerID uint   `gorm:"not null"` // Foreign key to QueueManager
	RouteID        uint   `gorm:"not null"` // Foreign key to Route
	Status         string `gorm:"size:20;not null;default:'active'"`
}
