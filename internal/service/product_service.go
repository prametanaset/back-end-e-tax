package service

import (
	"back-end-e-tax/internal/repository"
	"back-end-e-tax/pkg/model"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ProductService interface {
	GetProducts(tenantID uint) ([]model.Product, error)
	GetProductByID(id int, tenantID uint) (*model.Product, error)
	CreateProduct(product *model.Product, tenantID uint) error
	UpdateProduct(id int, input *model.Product, tenantID uint) (*model.Product, error)
	DeleteProduct(id int, tenantID uint) error
	GetPaginatedProducts(page, limit int, search string, tenantID uint) ([]model.Product, int64, error)
}

type productService struct {
	repo repository.ProductRepository
}

// Constructor ให้เรียกจาก main หรือ controller ได้
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProducts(tenantID uint) ([]model.Product, error) {
	return s.repo.GetAll(tenantID)
}

func (s *productService) GetProductByID(id int, tenantID uint) (*model.Product, error) {
	return s.repo.GetByID(id, tenantID)
}

func (s *productService) CreateProduct(product *model.Product, tenantID uint) error {
	product.TenantID = tenantID

	// สร้างรหัสสินค้าอัตโนมัติ
	lastProduct, err := s.repo.GetLastProductCode(tenantID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	lastNum := 0
	if lastProduct != nil && lastProduct.ProductCode != "" {
		_, err := fmt.Sscanf(lastProduct.ProductCode, "PROD-%04d", &lastNum)
		if err != nil {
			return err
		}
	}

	product.ProductCode = fmt.Sprintf("PROD-%04d", lastNum+1)

	return s.repo.Create(product)
}

func (s *productService) UpdateProduct(id int, input *model.Product, tenantID uint) (*model.Product, error) {
	product, err := s.repo.GetByID(id, tenantID)
	if err != nil {
		return nil, err
	}

	product.Name = input.Name
	product.PriceStandard = input.PriceStandard
	product.VatCategoryID = input.VatCategoryID
	product.UpdatedAt = time.Now()

	err = s.repo.Update(product)
	return product, err
}

func (s *productService) DeleteProduct(id int, tenantID uint) error {
	product, err := s.repo.GetByID(id, tenantID)
	if err != nil {
		return err
	}
	return s.repo.Delete(product)
}

func (s *productService) GetPaginatedProducts(page, limit int, search string, tenantID uint) ([]model.Product, int64, error) {
	return s.repo.GetPaginated(page, limit, search, tenantID)
}
