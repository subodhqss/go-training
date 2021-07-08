package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type productService interface {
	PrintProduct() []*models.Product
	PrintProductCode(string) *models.Product
	SaveProduct(models.Product) *models.Product
	EditProduct(models.Product) *models.Product
}

type proServ struct {
	proRepo repository.ProductReposiotry
}

func NewProductService(proRepo repository.ProductReposiotry) productService {
	return &proServ{proRepo: proRepo}
}

func (es *proServ) PrintProduct() []*models.Product {
	cus := es.proRepo.PrintProduct()
	return cus
}

func (es *proServ) PrintProductCode(eid string) *models.Product {
	cus := es.proRepo.PrintProductCode(eid)
	return cus
}

func (cus *proServ) SaveProduct(Product models.Product) *models.Product {
	cusId := cus.proRepo.SaveProduct(Product)
	return cusId
}
func (cus *proServ) EditProduct(Product models.Product) *models.Product {
	cusId := cus.proRepo.EditProduct(Product)
	return cusId
}