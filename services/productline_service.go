package services

import (
	// "internal/poll"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type ProductlineService interface {
	PrintProductline() []*models.Productline
	 SaveProductline(models.Productline) *models.Productline
	 UpdateProductline(models.Productline) *models.Productline
	// UpdatePl(models.Productline) *models.Productline
}

type prodliServ struct {
	prodliRepo repository.ProductlineReposiotry
}

func NewProductlineService(prodliRepo repository.ProductlineReposiotry) ProductlineService {
	return &prodliServ{prodliRepo: prodliRepo}
}

func (pls *prodliServ) PrintProductline() []*models.Productline {
	prodli := pls.prodliRepo.PrintProductline()
	return prodli
}

func (prodli *prodliServ) SaveProductline(Productline models.Productline) *models.Productline {

	prodliNo := prodli.prodliRepo.SaveProductline(Productline)
	return prodliNo
}


func (pls *prodliServ) UpdateProductline(Productline models.Productline) *models.Productline {
	prodCo := pls.prodliRepo.UpdateProductline(Productline)
	return prodCo
}
