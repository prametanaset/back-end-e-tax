// internal/dto/create_user_dto.go
package dto

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}
