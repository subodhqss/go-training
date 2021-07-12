package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type OrderService interface {
	PrintOrder() []*models.Order
	//PrintOrderId(string) []*models.Order
	SaveOrder(models.Order) *models.Order
	UpdateOrder(models.Order) *models.Order
	UpdateO(models.Order) *models.Order
}

type ordServ struct {
	ordRepo repository.OrderReposiotry
}

func NewOrderService(ordRepo repository.OrderReposiotry) OrderService{
	return &ordServ{ordRepo: ordRepo}
}

//Order model functions
func (ov *ordServ) PrintOrder() []*models.Order {
	ord := ov.ordRepo.PrintOrder()
	return ord
}
/*func (ov *orderServ) PrintOrderId(oid string) *models.Office {
	order := ov.ordRepo.PrintOrderId(oid)
	return order
}*/




func (ord *ordServ) SaveOrder(Order models.Order) *models.Order {

	ordNo := ord.ordRepo.SaveOrder(Order)
	return ordNo
}

func (ov *ordServ) UpdateOrder(Order models.Order) *models.Order {
	ordNo := ov.ordRepo.UpdateOrder(Order)
	return ordNo
}

func (ov *ordServ) UpdateO(Order models.Order) *models.Order {
	ordNo := ov.ordRepo.UpdateO(Order)
	return ordNo
}
