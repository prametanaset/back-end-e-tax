package model

import "time"

type AuthToken struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User
	Token     string
	ExpiresAt time.Time
	Revoked   bool
	IPAddress string
	UserAgent string
	CreatedAt time.Time
}
