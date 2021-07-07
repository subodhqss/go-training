package repository

import (
	
	"fmt"
	"log"
	"strconv"

	"github.com/subodhqss/go-training/models"
)

type CustomerReposiotry interface {
	PrintCustomer() []*models.Customer
	PrintCustomerId(string) *models.Customer
	SaveCustomer(models.Customer) *models.Customer
	UpdateCustomer(models.Customer) *models.Customer

}

func NewCusRepo() CustomerReposiotry {
	return &cusRepo{}
}

type cusRepo struct{}

func (er *cusRepo) PrintCustomer() []*models.Customer {
	var Customer []*models.Customer
	result := gormDB.Limit(10).Find(&Customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(result)
	return Customer
}

func (er *cusRepo) PrintCustomerId(code string) *models.Customer {

	c_id, _ := strconv.ParseInt(code, 0, 64) //type conversion
	Customer := &models.Customer{} 
	result := gormDB.Where("CustomerNumber", c_id).Find(Customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method", result, code)
	return Customer
}

func (tr *cusRepo) SaveCustomer(Customer models.Customer) *models.Customer {

	result := gormDB.Create(&Customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", Customer)
	return &Customer
}

func (tr *cusRepo) UpdateCustomer(Customer models.Customer) *models.Customer {
	result := gormDB.Model(&Customer).Where("CustomerNumber", Customer.CustomerNumber).Updates(Customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", Customer)
	return &Customer
}

