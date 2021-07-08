package repository

import (
	"fmt"
	"log"
	// "strconv"

	"github.com/subodhqss/go-training/models"
)

type ProductReposiotry interface {
	PrintProduct() []*models.Product
	PrintProductId(string) *models.Product
	SaveProduct(models.Product) *models.Product
	UpdateProduct(models.Product) *models.Product

}

func NewProRepo() ProductReposiotry {
	return &proRepo{}

}

type proRepo struct{}

func (er *proRepo) PrintProduct() []*models.Product {
	var Product []*models.Product
	result := gormDB.Limit(10).Find(&Product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return Product
}

func (er *proRepo) PrintProductId(code string) *models.Product {
	Product := &models.Product{} 
	result := gormDB.Where("ProductCode", code).Find(Product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method", result, code)
	return Product
}

func (tr *proRepo) SaveProduct(Product models.Product) *models.Product {

	result := gormDB.Create(&Product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", Product)
	return &Product
}

func (tr *proRepo) UpdateProduct(Product models.Product) *models.Product {
	result := gormDB.Model(&Product).Where("ProductCode", Product.ProductCode).Updates(Product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", Product)
	return &Product
}

