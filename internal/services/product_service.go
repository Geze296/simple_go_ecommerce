package services

import (
	"github.com/geze296/simple_go_ecommerce/internal/models"
	"github.com/geze296/simple_go_ecommerce/internal/repositories"
)

type ProductService struct{
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository)  *ProductService{
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) CreateProduct(product *models.Product) error{
	err := s.productRepository.Create(product)
	if err != nil {
		return err
	}
	return nil
}