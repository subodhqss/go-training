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
	EditCustomer(models.Customer) *models.Customer
}

func NewCustmRepo() CustomerReposiotry {
	return &custmRepo{}

}

type custmRepo struct{}

func (er *custmRepo) PrintCustomer() []*models.Customer {

	var customer []*models.Customer
	result := gormDB.Limit(10).Find(&customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	return customer
}

func (er *custmRepo) PrintCustomerId(cid string) *models.Customer {
	var customer *models.Customer
	c_id, _ := strconv.ParseInt(cid, 0, 64)

	result := gormDB.Preload("Customer").Where("customerNumber", c_id).Find(&customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(&customer)
	return customer
}

func (tr *custmRepo) SaveCustomer(customer models.Customer) *models.Customer {
	b := gormDB.First(&customer)
	result := gormDB.Save(&b)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println(&customer)
	return &customer

}

func (tr *custmRepo) EditCustomer(customer models.Customer) *models.Customer {

	result := gormDB.Model(&customer).Where("customerNumber", customer.CustomerNumber).Updates(customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println(&customer)
	return &customer
}
