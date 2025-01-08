package models

import "time"

type UserSession struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Token     string    `gorm:"size:512" json:"token"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	IssuedAt  time.Time `gorm:"autoCreateTime" json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
