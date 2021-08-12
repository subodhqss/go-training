package repository

import (
	 "fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type ProductReposiotry interface {
	 PrintProduct() []*models.Product
	 PrintProductCode(string) *models.Product
	// SaveProduct(models.Product) *models.Product
	// UpdateProduct(models.Product) *models.Product
	// UpdateP(models.Product) *models.Product
}

func NewProdRepo() ProductReposiotry {

	return &prodRepo{}
}

type prodRepo struct{}

func (pr *prodRepo) PrintProduct() []*models.Product {
	var product []*models.Product
	result := gormDB.Preload("ProductlineDetails").Limit(10).Find(&product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return product
}


func (pr *prodRepo) PrintProductCode(code string) *models.Product {
	var Product *models.Product

	result := gormDB.Preload("ProductlineDetails").Where("ProductCode", code).Find(&Product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method")
	return Product
}
