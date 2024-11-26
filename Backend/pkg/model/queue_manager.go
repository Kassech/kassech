package models

import "gorm.io/gorm"

type QueueManager struct {
    gorm.Model
    UserID uint
    Status string
}
