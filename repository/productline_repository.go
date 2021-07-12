package repository

import (
	
	"fmt"
	"log"
	"github.com/subodhqss/go-training/models"
)

type ProductlineReposiotry interface {
	PrintProductline() []*models.Productline
	PrintProductlineId(string) []*models.Productline
	SaveProductline(models.Productline) *models.Productline
	UpdateProductline(models.Productline) *models.Productline

}

func NewPrlRepo() ProductlineReposiotry {
	return &prlRepo{}
}

type prlRepo struct{}

func (er *prlRepo) PrintProductline() []*models.Productline {
	var Productline []*models.Productline
	result := gormDB.Limit(10).Find(&Productline)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(result)
	return Productline
}

func (er *prlRepo) PrintProductlineId(code string) []*models.Productline {

	
	var Productline []*models.Productline
	result := gormDB.Where("productLine", code).Find(&Productline)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method", result, code)
	return Productline
}

func (tr *prlRepo) SaveProductline(Productline models.Productline) *models.Productline {

	result := gormDB.Create(&Productline)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", Productline)
	return &Productline
}

func (tr *prlRepo) UpdateProductline(Productline models.Productline) *models.Productline {
	result := gormDB.Model(&Productline).Where("productLine", Productline.ProductLine).Updates(Productline)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", Productline)
	return &Productline

}