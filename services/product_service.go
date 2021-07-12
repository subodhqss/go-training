package service

import (
	// "internal/poll"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type ProductService interface {
	PrintProduct() []*models.Product
	PrintProductId(string) *models.Product
	SaveProduct(models.Product) *models.Product
	UpdateProduct(models.Product) *models.Product
	UpdateP(models.Product) *models.Product
}

type prodServ struct {
	prodRepo repository.ProductReposiotry
}

func NewProductService(prodRepo repository.ProductReposiotry) ProductService {
	return &prodServ{prodRepo: prodRepo}
}

//Product model functions
func (ps *prodServ) PrintProduct() []*models.Product {
	prod := ps.prodRepo.PrintProduct()
	return prod
}
func (ps *prodServ) PrintProductId(code string) *models.Product {
	prod := ps.prodRepo.PrintProductId(code)
	return prod
}



func (prod *prodServ) SaveProduct(Product models.Product) *models.Product {

	custNo := prod.prodRepo.SaveProduct(Product)
	return custNo
}

func (ps *prodServ) UpdateProduct(Product models.Product) *models.Product {
	prodCo := ps.prodRepo.UpdateProduct(Product)
	return prodCo
}

func (ps *prodServ) UpdateP(Product models.Product) *models.Product {
	prodCo := ps.prodRepo.UpdateP(Product)
	return prodCo
}
