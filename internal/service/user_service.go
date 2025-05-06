// pkg/service/user_service.go
package service

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"back-end-e-tax/internal/repository"
	"back-end-e-tax/pkg/model"
)

type UserService interface {
	GetAll(tenantID uint) ([]model.User, error)
	GetByID(id, tenantID uint) (*model.User, error)
	Register(user *model.User) error
	Login(username, password string, tenantID uint) (*model.User, error)
	Update(user *model.User) error
	Delete(id, tenantID uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAll(tenantID uint) ([]model.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var filtered []model.User
	for _, u := range users {
		if u.TenantID == tenantID {
			filtered = append(filtered, u)
		}
	}
	return filtered, nil
}

func (s *userService) GetByID(id, tenantID uint) (*model.User, error) {
	u, err := s.repo.GetByID(id)
	if err != nil || u.TenantID != tenantID {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (s *userService) Register(user *model.User) error {
	// check duplicate username
	if existing, err := s.repo.FindByUsername(user.Username); err == nil {
		if existing.TenantID == user.TenantID {
			return errors.New("username already exists")
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// check duplicate email
	if existing, err := s.repo.FindByEmail(user.Email); err == nil {
		if existing.TenantID == user.TenantID {
			return errors.New("email already registered")
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return s.repo.Create(user)
}

func (s *userService) Login(username, password string, tenantID uint) (*model.User, error) {
	u, err := s.repo.FindByUsername(username)
	if err != nil || u.TenantID != tenantID {
		return nil, errors.New("invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}

	u.LastLogin = time.Now()
	if err := s.repo.Update(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *userService) Update(user *model.User) error {
	// ensure the user belongs to the tenant
	existing, err := s.repo.GetByID(user.ID)
	if err != nil || existing.TenantID != user.TenantID {
		return errors.New("user not found")
	}
	// re-hash password if provided
	if user.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashed)
	} else {
		user.Password = existing.Password
	}
	return s.repo.Update(user)
}

func (s *userService) Delete(id, tenantID uint) error {
	u, err := s.repo.GetByID(id)
	if err != nil || u.TenantID != tenantID {
		return errors.New("user not found")
	}
	return s.repo.Delete(u)
}
