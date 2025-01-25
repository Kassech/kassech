// models/queue_manager_route.go
package models

import "gorm.io/gorm"

type QueueManagerRoute struct {
	gorm.Model
	UserID    uint    `gorm:"not null"`
	StationID uint    `gorm:"not null"`
	Paths     []Path  `gorm:"many2many:queue_manager_route_paths;"`
	User      User    `gorm:"foreignKey:UserID"`
	Station   Station `gorm:"foreignKey:StationID"`
}
