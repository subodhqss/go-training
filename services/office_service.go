package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type officeService interface {
	PrintOffice() []*models.Office
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
