package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type CustomerService interface {
	PrintCustomer() []*models.Customer
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
