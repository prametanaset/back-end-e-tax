package models

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	StoreId     uint      `gorm:"type:integer;not null"`
	ProductCode string 		`gorm:"type:varchar(100);not null"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	Vat       	bool  	`gorm:"not null"`
	VatRate     float64   `gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
