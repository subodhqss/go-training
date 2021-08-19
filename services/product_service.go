package service

import (
	"fmt"
	"mime/multipart"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/util"
)

type ProductService interface {
	PrintProduct() []*models.Product
	PrintProductId(string) *models.Product
	SaveProduct(models.Product) *models.Product
	UpdateProduct(models.Product) *models.Product
	UpdateImage(string, multipart.File, *multipart.FileHeader, string)
}

type proServ struct {
	proRepo repository.ProductReposiotry
}

func NewProductService(proRepo repository.ProductReposiotry) ProductService {
	return &proServ{proRepo: proRepo}
}

func (ps *proServ) PrintProduct() []*models.Product {
	pro := ps.proRepo.PrintProduct()
	return pro
}

func (ps *proServ) PrintProductId(code string) *models.Product {
	pro := ps.proRepo.PrintProductId(code)
	return pro
}

func (ps *proServ) SaveProduct(Product models.Product) *models.Product {

	pro := ps.proRepo.SaveProduct(Product)
	return pro
}

func (ps *proServ) UpdateProduct(Product models.Product) *models.Product {
	pro := ps.proRepo.UpdateProduct(Product)
	return pro
}

func (ps *proServ) UpdateImage(code1 string, file multipart.File, header *multipart.FileHeader, host string) {

	fmt.Println("serv")
	imagePath := util.UploadFile(file, header, code1)

	ps.proRepo.UpdateImage(host, imagePath, code1)

}
