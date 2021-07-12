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


func (pr *prlServ) PrintProductline() []*models.Productline {
	prol := pr.prlRepo.PrintProductline()
	return prol
}

func (pr *prlServ) PrintProductlineId(code string) []*models.Productline {
	prol := pr.prlRepo.PrintProductlineId(code)
	return prol
}

func (pr *prlServ) SaveProductline(Productline models.Productline) *models.Productline {

	prol := pr.prlRepo.SaveProductline(Productline)
	return prol
}

func (pl *prlServ) UpdateProductline(Productline models.Productline) *models.Productline {
	prol := pl.prlRepo.UpdateProductline(Productline)
	return prol
}
