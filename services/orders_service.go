package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type OrderService interface {
	PrintOrder() []*models.Order
	PrintOrderId(string) []*models.Order
	SaveOrder(models.Order) *models.Order
	UpdateOrder(models.Order) *models.Order
}

type ordServ struct {
	ordRepo repository.OrderReposiotry
}

func NewOrderService(ordRepo repository.OrderReposiotry) OrderService {
	return &ordServ{ordRepo: ordRepo}
}


func (ors *ordServ) PrintOrder() []*models.Order {
	ord := ors.ordRepo.PrintOrder()
	return ord
}

func (ors *ordServ) PrintOrderId(code string) []*models.Order {
	ord := ors.ordRepo.PrintOrderId(code)
	return ord
}

func (ors *ordServ) SaveOrder(Order models.Order) *models.Order {

	ord := ors.ordRepo.SaveOrder(Order)
	return ord
}

func (ors *ordServ) UpdateOrder(Order models.Order) *models.Order {
	ord := ors.ordRepo.UpdateOrder(Order)
	return ord
}
