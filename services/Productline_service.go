package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type ProductlineService interface {
	PrintProductline() []*models.Productline
	PrintProductlineId(string) []*models.Productline
	SaveProductline(models.Productline) *models.Productline
	UpdateProductline(models.Productline) *models.Productline
}

type prlServ struct {
	prlRepo repository.ProductlineReposiotry
}

func NewProductlineService(prlRepo repository.ProductlineReposiotry) ProductlineService {
	return &prlServ{prlRepo: prlRepo}
}


func (pys *prlServ) PrintProductline() []*models.Productline {
	prl := pys.prlRepo.PrintProductline()
	return prl
}

func (pys *prlServ) PrintProductlineId(code string) []*models.Productline {
	prl := pys.prlRepo.PrintProductlineId(code)
	return prl
}

func (pys *prlServ) SaveProductline(Productline models.Productline) *models.Productline {

	prl := pys.prlRepo.SaveProductline(Productline)
	return prl
}

func (pys *prlServ) UpdateProductline(Productline models.Productline) *models.Productline {
	prl := pys.prlRepo.UpdateProductline(Productline)
	return prl
}
