package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type OfficeRepository interface {
	PrintOffice() []*models.Office
	PrintOfficeId(string) *models.Office
	SaveOffice(models.Office) *models.Office
	EditOffice(models.Office) *models.Office
}

func NewOfficeRepo() OfficeRepository {
	return &officeRepo{}
}

type officeRepo struct{}

func (or *officeRepo) PrintOffice() []*models.Office {
	var office []*models.Office
	result := gormDB.Limit(10).Find(&office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return office
}

func (er *officeRepo) PrintOfficeId(code string) *models.Office {
	var office *models.Office

	result := gormDB.Preload("Employees").Where("officeCode", code).Find(&office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(&office)
	return office
}

func (tr *officeRepo) SaveOffice(office models.Office) *models.Office {
	b := gormDB.First(&office)
	result := gormDB.Save(&b)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println(&office)
	return &office

}

func (tr *officeRepo) EditOffice(office models.Office) *models.Office {

	result := gormDB.Model(&office).Where("officeCode", office.OfficeCode).Updates(office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println(&office)
	return &office
}
