// internal/service/user_service.go
package service

import (
	"back-end-e-tax/internal/repository"
	"back-end-e-tax/pkg/model"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) CreateUser(user *model.User) error {
	// ตรวจสอบว่ามี tenant จริง
	if user.TenantID == 0 {
		return fmt.Errorf("tenant_id is required")
	}

	// Optional: เช็ค username ซ้ำ
	existing, err := s.repo.GetByUsername(user.Username)
	if err == nil && existing != nil {
		return fmt.Errorf("username already exists")
	}

	// ใส่เวลา login ล่าสุด
	if user.LastLogin.IsZero() {
		user.LastLogin = time.Now()
	}

	// ถ้ายังไม่ได้ hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password failed")
	}
	user.Password = string(hashed)

	// บันทึก
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(user)
}
