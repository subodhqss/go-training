package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type ProductReposiotry interface {
	PrintProduct() []*models.Product
	SaveProduct(models.Product) *models.Product
	UpdateProduct(models.Product) *models.Product
	UpdateP(models.Product) *models.Product

}
func NewProdRepo() ProductReposiotry {
	
	return &prodRepo{}


}

type prodRepo struct{}

//Product model functions
func (pr *prodRepo) PrintProduct() []*models.Product {
	var product []*models.Product
	result := gormDB.Limit(10).Find(&product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return product
}

func (pr *prodRepo) SaveProduct(product models.Product) *models.Product {

	result := gormDB.Create(&product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", product)
	return &product
}

func (pr *prodRepo) UpdateProduct(product models.Product) *models.Product {
	result := gormDB.Model(&product).Where("Code", product.ProductCode).Updates(product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", product)
	return &product
}

func (pr *prodRepo) UpdateP(product models.Product) *models.Product {
	result := gormDB.Model(&product).Where("productCode", product.ProductCode).Updates(product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", product)
	return &product
}







