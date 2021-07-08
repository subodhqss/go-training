package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/subodhqss/go-training/models"
)

type ProductReposiotry interface {
	PrintProduct() []*models.Product
	PrintProductCode(string) *models.Product
	SaveProduct(models.Product) *models.Product
	EditProduct(models.Product) *models.Product
}

func NewProRepo() ProductReposiotry {
	return &proRepo{}

}

type proRepo struct{}

func (er *proRepo) PrintProduct() []*models.Product {

	var product []*models.Product
	result := gormDB.Limit(10).Find(&product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	return product
}

func (er *proRepo) PrintProductCode(code string) *models.Product {
	var product *models.Product
	c_id, _ := strconv.ParseInt(code, 0, 64)

	result := gormDB.Preload("Product").Where("productNumber", c_id).Find(&product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(&product)
	return product
}

func (tr *proRepo) SaveProduct(product models.Product) *models.Product {
	// b := gormDB.First(&customer)
	result := gormDB.First(&product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println(&product)
	return &product

}

func (tr *proRepo) EditProduct(product models.Product) *models.Product {

	result := gormDB.Model(&product).Where("productNumber", product.ProductCode).Updates(product)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println(&product)
	return &product
}
