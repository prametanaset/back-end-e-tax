package models

import (
	"gorm.io/gorm"
)

// Customer คำอธิบายของลูกค้า
type CustomerResponse struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
