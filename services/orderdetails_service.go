package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type OrderDetailService interface {
	PrintOrderDetail() []*models.Orderdetail
	PrintOrderDetailId(string) *models.Orderdetail
	SaveOrderDetail(models.Orderdetail) *models.Orderdetail
	UpdateOrderDetail(models.Orderdetail) *models.Orderdetail
}

type orddServ struct {
	orddRepo repository.OrderDetailReposiotry
}

func NewOrderDetailService(orddRepo repository.OrderDetailReposiotry) OrderDetailService {
	return &orddServ{orddRepo: orddRepo}
}


func (cs *orddServ) PrintOrderDetail() []*models.Orderdetail {
	ordd := cs.orddRepo.PrintOrderDetail()
	return ordd
}

func (cs *orddServ) PrintOrderDetailId(code string) *models.Orderdetail {
	ordd := cs.orddRepo.PrintOrderDetailId(code)
	return ordd
}

func (cs *orddServ) SaveOrderDetail(OrderDetail models.Orderdetail) *models.Orderdetail {

	ordd := cs.orddRepo.SaveOrderDetail(OrderDetail)
	return ordd
}

func (cs *orddServ) UpdateOrderDetail(OrderDetail models.Orderdetail) *models.Orderdetail {
	ordd := cs.orddRepo.UpdateOrderDetail(OrderDetail)
	return ordd
}
