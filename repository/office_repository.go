package repository

import (
	"log"
	
	"github.com/subodhqss/go-training/models"
)
type OfficeRepository interface {
	PrintOffice() []*models.Office
}

func NewOfficeRepo() OfficeRepository {
	return &officeRepo{}
}

type officeRepo struct{}
func (or *officeRepo) PrintOffice() []*models.Office{
	var office []*models.Office
	result := gormDB.Limit(10).Find(&office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return office
}