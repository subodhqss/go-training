package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type OrderService interface {
	PrintOrder() []*models.Order
	PrintOrderId(string) *models.Order
	SaveOrder(models.Order) *models.Order
	UpdateOrder(models.Order) *models.Order
}

type ordServ struct {
	ordRepo repository.OrderReposiotry
}

func NewOrderService(ordRepo repository.OrderReposiotry) OrderService {
	return &ordServ{ordRepo: ordRepo}
}

func (od *ordServ) PrintOrder() []*models.Order {
	ordd := od.ordRepo.PrintOrder()
	return ordd
}

func (od *ordServ) PrintOrderId(code string) *models.Order {
	ord := od.ordRepo.PrintOrderId(code)
	return ord
}

func (od *ordServ) SaveOrder(Order models.Order) *models.Order {

	ord := od.ordRepo.SaveOrder(Order)
	return ord
}

func (od *ordServ) UpdateOrder(Order models.Order) *models.Order {
	ord := od.ordRepo.UpdateOrder(Order)
	return ord
}
