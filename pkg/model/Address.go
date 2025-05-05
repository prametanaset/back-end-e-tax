package model

type Address struct {
	ID         uint `gorm:"primaryKey"`
	Street     string
	District   string
	Province   string
	PostalCode string
	Country    string `gorm:"default:Thailand"`
}
