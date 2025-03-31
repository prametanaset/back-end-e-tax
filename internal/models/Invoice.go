package models

import "gorm.io/gorm"

type Invoice struct {
    gorm.Model
    InvoiceNumber string  `gorm:"unique;not null"`
    CustomerName  string  `gorm:"not null"`
    Amount        float64 `gorm:"not null"`
    Status        string  `gorm:"default:'Pending'"`
}
