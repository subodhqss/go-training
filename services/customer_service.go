package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type CustomerService interface {
	PrintCustomer() []*models.Customer
	SaveCustomer(models.Customer) *models.Customer
	UpdateCustomer(models.Customer) *models.Customer
	UpdateC(models.Customer) *models.Customer
}

type custServ struct {
	custRepo repository.CustomerReposiotry
}

func NewCustomerService(custRepo repository.CustomerReposiotry) CustomerService {
	return &custServ{custRepo: custRepo}
}

//Customer model functions
func (cs *custServ) PrintCustomer() []*models.Customer {
	cust := cs.custRepo.PrintCustomer()
	return cust
}

/*
func (cs *custServ) PrintCustomerNumber(number int) *models.Customer {
	cust := cs.custRepo.PrintCustomerNumber(number)
	return cust
}
*/
func (cust *custServ) SaveCustomer(Customer models.Customer) *models.Customer {

	custNo := cust.custRepo.SaveCustomer(Customer)
	return custNo
}

func (cs *custServ) UpdateCustomer(Customer models.Customer) *models.Customer {
	custNo := cs.custRepo.UpdateCustomer(Customer)
	return custNo
}

func (cs *custServ) UpdateC(Customer models.Customer) *models.Customer {
	custNo := cs.custRepo.UpdateC(Customer)
	return custNo
}
