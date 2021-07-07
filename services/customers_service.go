package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type customerService interface {
	PrintCustomer() []*models.Customer
	PrintCustomerId(string) *models.Customer
	SaveCustomer(models.Customer) *models.Customer
	EditCustomer(models.Customer) *models.Customer
}

type custmServ struct {
	custmRepo repository.CustomerReposiotry
}

func NewCustomerService(custmRepo repository.CustomerReposiotry) customerService {
	return &custmServ{custmRepo: custmRepo}
}

func (es *custmServ) PrintCustomer() []*models.Customer {
	cus := es.custmRepo.PrintCustomer()
	return cus
}

func (es *custmServ) PrintCustomerId(eid string) *models.Customer {
	cus := es.custmRepo.PrintCustomerId(eid)
	return cus
}

func (cus *custmServ) SaveCustomer(Customer models.Customer) *models.Customer {
	cusId := cus.custmRepo.SaveCustomer(Customer)
	return cusId
}
func (cus *custmServ) EditCustomer(Customer models.Customer) *models.Customer {
	cusId := cus.custmRepo.EditCustomer(Customer)
	return cusId
}