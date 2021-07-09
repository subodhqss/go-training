package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type ProductlineReposiotry interface {
	PrintProductline() []*models.Productline
	SaveProductline(models.Productline) *models.Productline
	UpdateProductline(models.Productline) *models.Productline
	UpdatePl(models.Productline) *models.Productline
}

func NewProdLineRepo() ProductlineReposiotry {

	return &prodliRepo{}

}

type prodliRepo struct{}

//ProductLine model functions
func (pr *prodliRepo) PrintProductline() []*models.Productline {
	var productline []*models.Productline
	result := gormDB.Limit(10).Find(&productline)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return productline
}

func (plr *prodliRepo) SaveProductline(productline models.Productline) *models.Productline {

	result := gormDB.Create(&productline)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", productline)
	return &productline
}

func (plr *prodliRepo) UpdateProductline(productline models.Productline) *models.Productline {
	result := gormDB.Model(&productline).Where("Code", productline.ProductLine).Updates(productline)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", productline)
	return &productline
}

func (plr *prodliRepo) UpdatePl(productline models.Productline) *models.Productline {
	result := gormDB.Model(&productline).Where("productCode", productline.ProductLine).Updates(productline)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", productline)
	return &productline
}
