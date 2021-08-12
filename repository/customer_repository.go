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
	UpdateC(models.Customer) *models.Customer
}

func NewCustRepo() CustomerReposiotry {

	return &custRepo{}

}
type custRepo struct{}

//Customer model functions
func (cr *custRepo) PrintCustomer() []*models.Customer {
	var customer []*models.Customer
	result := gormDB.Limit(10).Find(&customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return customer
}
func (cr *custRepo) PrintCustomerId(code string) *models.Customer {
	var customer *models.Customer

	c_id, _ := strconv.ParseInt(code, 0, 64) //type conversion

	result := gormDB.Preload("PaymentDetails").Where("CustomerNumber", c_id).Find(customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method")
	return customer
}

func (cr *custRepo) SaveCustomer(customer models.Customer) *models.Customer {

	result := gormDB.Create(&customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", customer)
	return &customer
}

func (cr *custRepo) UpdateCustomer(customer models.Customer) *models.Customer {
	result := gormDB.Model(&customer).Where("customerNumber", customer.CustomerNumber).Updates(customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", customer)
	return &customer
}

func (cr *custRepo) UpdateC(customer models.Customer) *models.Customer {
	result := gormDB.Model(&customer).Where("customerNumber", customer.CustomerNumber).Updates(customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", customer)
	return &customer
}

// func (es *custRepo) PrintCustomerNumber(eid string) *models.Employee {
// 	emp := es.custRepo.PrintEmployeId(eid)
// 	return emp
// }