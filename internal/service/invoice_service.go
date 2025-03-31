package service

import (
	"back-end-e-tax/internal/models"
)

func GetAllUsers() []models.User {
	return []models.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
	}
}