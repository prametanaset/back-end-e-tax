package model

import "time"

type TaxRate struct {
	ID        uint `gorm:"primaryKey"`
	Rate      float64
	StartDate time.Time
	EndDate   *time.Time
}
