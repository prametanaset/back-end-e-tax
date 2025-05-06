// pkg/repository/address_repository.go
package repository

import (
	"back-end-e-tax/pkg/model"

	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(addr *model.Address) error
}

type addressRepository struct{ db *gorm.DB }

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) Create(addr *model.Address) error {
	return r.db.Create(addr).Error
}
