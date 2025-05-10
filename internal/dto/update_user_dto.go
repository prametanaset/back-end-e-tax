package dto

type UpdateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}
