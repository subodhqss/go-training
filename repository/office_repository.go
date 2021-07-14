package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type OfficeReposiotry interface {
	PrintOffice() []*models.Office
	PrintOfficeId(string) *models.Office
	SaveOffice(models.Office) *models.Office
	EditOffice(models.Office) *models.Office
}

func NewOfcRepo() OfficeReposiotry {
	return &ofcRepo{}

}

type ofcRepo struct{}

func (er *ofcRepo) PrintOffice() []*models.Office {

	var office []*models.Office
	result := gormDB.Preload("Employees").Limit(10).Find(&office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	return office
}

func (er *ofcRepo) PrintOfficeId(code string) *models.Office {
	var office *models.Office

	result := gormDB.Preload("Employees").Where("officeCode", code).Find(&office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(&office)
	return office
}

func (tr *ofcRepo) SaveOffice(office models.Office) *models.Office {
	// b := gormDB.First(&office)
	result := gormDB.Create(&office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("created",office)
	return &office

}

func (tr *ofcRepo) EditOffice(office models.Office) *models.Office {

	result := gormDB.Model(&office).Where("officeCode", office.OfficeCode).Updates(office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("updated",office)
	return &office
}
