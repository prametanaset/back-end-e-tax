package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	TenantID  uint      `gorm:"not null" json:"tenant_id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	FullName  string    `json:"full_name"`
	Email     string    `gorm:"not null" json:"email"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	LastLogin time.Time `json:"last_login"`
	Roles     []Role    `gorm:"many2many:user_roles" json:"roles"`
}
