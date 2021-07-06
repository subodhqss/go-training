package repository

import (
	"fmt"
	"log"
	// "strconv"

	"github.com/subodhqss/go-training/models"
)

type OfficeReposiotry interface {
	PrintOffice() []*models.Office
	PrintOfficeId(string) *models.Office
	SaveOffice(models.Office) *models.Office
	UpdateOffice(models.Office) *models.Office

}

func NewOffRepo() OfficeReposiotry {
	return &offRepo{}

}

type offRepo struct{}

func (er *offRepo) PrintOffice() []*models.Office {
	var office []*models.Office
	result := gormDB.Limit(10).Find(&office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return office
}

func (er *offRepo) PrintOfficeId(code string) *models.Office {
	office := &models.Office{} 
	result := gormDB.Preload("Employees").Where("officeCode", code).Find(office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method", result, code)
	return office
}

func (tr *offRepo) SaveOffice(office models.Office) *models.Office {

	result := gormDB.Create(&office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", office)
	return &office
}

func (tr *offRepo) UpdateOffice(office models.Office) *models.Office {
	result := gormDB.Model(&office).Where("officeCode", office.OfficeCode).Updates(office)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", office)
	return &office
}

