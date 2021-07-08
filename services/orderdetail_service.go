package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type OrderdetailService interface {
	PrintOrderdetail() []*models.Orderdetail
	SaveDetail(models.Orderdetail) *models.Orderdetail
	UpdateDetail(models.Orderdetail) *models.Orderdetail
	UpdateD(models.Orderdetail) *models.Orderdetail
}
type orderServ struct {
	orderRepo repository.OrderdetailReposiotry
}

func NewOrderdetailService(orderRepo repository.OrderdetailReposiotry) OrderdetailService {
	return &orderServ{orderRepo: orderRepo}
}

//Orderdetail model functions
func (os *orderServ) PrintOrderdetail() []*models.Orderdetail {
	order := os.orderRepo.PrintOrderdetail()
	return order
}

/*
func (cs *custServ) PrintCustomerNumber(number int) *models.Customer {
	cust := cs.custRepo.PrintCustomerNumber(number)
	return cust
}
*/
func (order *orderServ) SaveDetail(Orderdetail models.Orderdetail) *models.Orderdetail {

	orderNo := order.orderRepo.SaveDetail(Orderdetail)
	return orderNo
}

func (os *orderServ) UpdateDetail(Orderdetail models.Orderdetail) *models.Orderdetail {
	orderNo := os.orderRepo.UpdateDetail(Orderdetail)
	return orderNo
}

func (os *orderServ) UpdateD(Orderdetail models.Orderdetail) *models.Orderdetail {
	orderNo := os.orderRepo.UpdateD(Orderdetail)
	return orderNo
}
