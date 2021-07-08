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

func (pro *proServ) PrintProduct() []*models.Product {
	pr := pro.proRepo.PrintProduct()
	return pr
}

func (pro *proServ) PrintProductCode(eid string) *models.Product {
	pr := pro.proRepo.PrintProductCode(eid)
	return pr
}

func (pro *proServ) SaveProduct(Product models.Product) *models.Product {
	pcd := pro.proRepo.SaveProduct(Product)
	return pcd
}
func (pro *proServ) EditProduct(Product models.Product) *models.Product {
	pcd := pro.proRepo.EditProduct(Product)
	return pcd
}