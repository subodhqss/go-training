package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type officeService interface {
	PrintOffice() []*models.Office
	PrintOfficeId(string) *models.Office
	SaveOffice(models.Office) *models.Office
	UpdateOffice(models.Office) *models.Office
	

	
}

type offServ struct {
	offRepo repository.OfficeReposiotry
}

func NewOfficeService(offRepo repository.OfficeReposiotry) officeService {
	return &offServ{offRepo: offRepo}
}


func (es *offServ) PrintOffice() []*models.Office {
	emp := es.offRepo.PrintOffice()
	return emp
}

func (es *offServ) PrintOfficeId(code string) *models.Office {
	off := es.offRepo.PrintOfficeId(code)
	return off
}

func (es *offServ) SaveOffice(Office models.Office) *models.Office {

	off := es.offRepo.SaveOffice(Office)
	return off
}

func (em *offServ) UpdateOffice(Office models.Office) *models.Office {
	off := em.offRepo.UpdateOffice(Office)
	return off
}
