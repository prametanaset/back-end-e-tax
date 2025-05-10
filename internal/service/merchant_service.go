// pkg/service/merchant_service.go
package service

import (
	"errors"

	"back-end-e-tax/internal/dto"
	"back-end-e-tax/internal/repository"
	"back-end-e-tax/pkg/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MerchantService interface {
	RegisterMerchant(
		legalName, branchCode, taxID string,
		addrDTO dto.AddressDTO,
		logoURL string,
		username, password, fullName, email string,
	) (uint, error)
}

type merchantService struct {
	addrRepo   repository.AddressRepository
	tenantRepo repository.TenantRepository
	userRepo   repository.UserRepository
	db         *gorm.DB
}

func NewMerchantService(
	ar repository.AddressRepository,
	tr repository.TenantRepository,
	ur repository.UserRepository,
	db *gorm.DB,
) MerchantService {
	return &merchantService{ar, tr, ur, db}
}

func (s *merchantService) RegisterMerchant(
	legalName, branchCode, taxID string,
	addrDTO dto.AddressDTO,
	logoURL string,
	username, password, fullName, email string,
) (uint, error) {
	var tenantID uint

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// 1) create Address
		addr := &model.Address{
			Street:     addrDTO.Street,
			District:   addrDTO.District,
			Province:   addrDTO.Province,
			PostalCode: addrDTO.PostalCode,
			Country:    addrDTO.Country,
		}
		if err := tx.Create(addr).Error; err != nil {
			return err
		}

		// 2) create Tenant
		tenant := &model.Tenant{
			LegalName:  legalName,
			BranchCode: branchCode,
			TaxID:      taxID,
			AddressID:  addr.ID,
			LogoURL:    logoURL,
		}
		if err := tx.Create(tenant).Error; err != nil {
			return err
		}

		// 3) check duplicate username/email
		if _, err := s.userRepo.FindByUsername(username); err == nil {
			return errors.New("username already taken")
		}
		if _, err := s.userRepo.FindByEmail(email); err == nil {
			return errors.New("email already registered")
		}

		// 4) hash password & create User
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user := &model.User{
			TenantID: tenant.ID,
			Username: username,
			Password: string(hashed),
			FullName: fullName,
			Email:    email,
			IsActive: true,
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		tenantID = tenant.ID
		return nil
	})

	return tenantID, err
}
