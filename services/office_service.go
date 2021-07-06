package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type officeService interface {
	PrintOffice() []*models.Office
	PrintOfficeId(string) *models.Office
	SaveOffice(models.Office) *models.Office
	EditOffice(models.Office) *models.Office
}

type ofcServ struct {
	ofcRepo repository.OfficeReposiotry
}

func NewOfficeService(ofcRepo repository.OfficeReposiotry) officeService {
	return &ofcServ{ofcRepo: ofcRepo}
}

func (es *ofcServ) PrintOffice() []*models.Office {
	ofc := es.ofcRepo.PrintOffice()
	return ofc
}

func (es *ofcServ) PrintOfficeId(code string) *models.Office {
	ofc := es.ofcRepo.PrintOfficeId(code)
	return ofc
}

func (ofc *ofcServ) SaveOffice(Office models.Office) *models.Office {
	ofcId := ofc.ofcRepo.SaveOffice(Office)
	return ofcId
}
func (ofc *ofcServ) EditOffice(Office models.Office) *models.Office {
	ofcId := ofc.ofcRepo.EditOffice(Office)
	return ofcId
}
