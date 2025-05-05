package model

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
