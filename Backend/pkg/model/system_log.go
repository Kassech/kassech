package models

import (
	"time"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type SystemLog struct {
	gorm.Model
	EventType  string         `gorm:"not null"`
	EntityType string         `gorm:"not null"`
	EntityID   uint           `gorm:"not null"`
	Details    pgtype.JSONB   `gorm:"not null;type:jsonb"`
	CreatedAt  time.Time      `gorm:"default:now()"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
