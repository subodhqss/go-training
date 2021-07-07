package repository

import (
	"log"

	"github.com/subodhqss/go-training/models"
)

type CustomerReposiotry interface {
	PrintCustomer() []*models.Customer
}
func NewCustRepo() CustomerReposiotry {
	
	return &custRepo{}

}

type custRepo struct{}

//Employee model functions
func (cr *custRepo) PrintCustomer() []*models.Customer {
	var customer []*models.Customer
	result := gormDB.Limit(10).Find(&customer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return customer
}

