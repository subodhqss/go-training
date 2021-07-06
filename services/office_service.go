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

type officeServ struct {
	officeRepo repository.OfficeRepository
}

func NewOfficeService(officeRepo repository.OfficeRepository) officeService {
	return &officeServ{officeRepo: officeRepo}

}
func (os *officeServ) PrintOffice() []*models.Office {
	office := os.officeRepo.PrintOffice()
	return office
}


func (os *officeServ) PrintOfficeId(code string) *models.Office {
	office := os.officeRepo.PrintOfficeId(code)
	return office
}

func (os *officeServ) SaveOffice(Office models.Office) *models.Office {
	officeId := os.officeRepo.SaveOffice(Office)
	return officeId
}

func (os *officeServ) EditOffice(Office models.Office) *models.Office {
	officeId := os.officeRepo.EditOffice(Office)
	return officeId
}
