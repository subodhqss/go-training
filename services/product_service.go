package services

import (
	// "internal/poll"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type ProductService interface {
	PrintProduct() []*models.Product
	 PrintProductCode(string) *models.Product
	// SaveProduct(models.Product) *models.Product
	// UpdateProduct(models.Product) *models.Product
	// UpdateP(models.Product) *models.Product
}

type prodServ struct {
	prodRepo repository.ProductReposiotry
}

func NewProductService(prodRepo repository.ProductReposiotry) ProductService {
	return &prodServ{prodRepo: prodRepo}
}

func (ps *prodServ) PrintProduct() []*models.Product {
	prod := ps.prodRepo.PrintProduct()
	return prod
}



func (ps *prodServ) PrintProductCode(code string) *models.Product {
	prod := ps.prodRepo.PrintProductCode(code)
	return prod
}