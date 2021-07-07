package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type customerService interface {
	PrintCustomer() []*models.Customer
	PrintCustomerId(string) *models.Customer
	SaveCustomer(models.Customer) *models.Customer
	UpdateCustomer(models.Customer) *models.Customer
}

type cusServ struct {
	cusRepo repository.CustomerReposiotry
}

func NewCustomerService(cusRepo repository.CustomerReposiotry) customerService {
	return &cusServ{cusRepo: cusRepo}
}


func (cs *cusServ) PrintCustomer() []*models.Customer {
	cus := cs.cusRepo.PrintCustomer()
	return cus
}

func (cs *cusServ) PrintCustomerId(code string) *models.Customer {
	cus := cs.cusRepo.PrintCustomerId(code)
	return cus
}

func (cs *cusServ) SaveCustomer(Customer models.Customer) *models.Customer {

	cus := cs.cusRepo.SaveCustomer(Customer)
	return cus
}

func (cs *cusServ) UpdateCustomer(Customer models.Customer) *models.Customer {
	cus := cs.cusRepo.UpdateCustomer(Customer)
	return cus
}
