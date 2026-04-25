package service

import (
	"errors"
	"prodexa/internal/models"
	"prodexa/internal/repository"
)

type ProductService interface {
	CreateProduct(p *models.Product, tenantID uint) error
	GetProducts(tenantID uint) ([]models.Product, error)
	GetProduct(id, tenantID uint) (*models.Product, error)
	UpdateProduct(p *models.Product, tenantID uint) error
	DeleteProduct(id, tenantID uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{repo: r}
}

// CREATE
func (s *productService) CreateProduct(p *models.Product, tenantID uint) error {
	p.TenantID = tenantID
	return s.repo.Create(p)
}

// READ ALL
func (s *productService) GetProducts(tenantID uint) ([]models.Product, error) {
	return s.repo.GetAll(tenantID)
}

// READ ONE
func (s *productService) GetProduct(id, tenantID uint) (*models.Product, error) {
	return s.repo.GetByID(id, tenantID)
}

// UPDATE (IMPORTANT CHECK ❗)
func (s *productService) UpdateProduct(p *models.Product, tenantID uint) error {
	existing, err := s.repo.GetByID(p.ID, tenantID)
	if err != nil {
		return errors.New("product not found")
	}

	p.TenantID = existing.TenantID
	return s.repo.Update(p)
}

// DELETE
func (s *productService) DeleteProduct(id, tenantID uint) error {
	return s.repo.Delete(id, tenantID)
}
