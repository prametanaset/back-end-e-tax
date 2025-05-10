package model

type VatCategory struct {
	ID          uint `gorm:"primaryKey"`
	Description string
	VatRate     float64
}
