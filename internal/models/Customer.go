package models

import (
	"time"
)

type Customer struct {
	ID        uint      `gorm:"primaryKey"`
	FirstName string    `gorm:"type:varchar(100);not null"`
	LastName  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Phone     string    `gorm:"type:varchar(15)"`
	Address  string    	`gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
