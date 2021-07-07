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

func (od *orddServ) PrintOrderDetail() []*models.Orderdetail {
	ordd := od.orddRepo.PrintOrderDetail()
	return ordd
}

func (od *orddServ) PrintOrderDetailId(code string) *models.Orderdetail {
	ord := od.orddRepo.PrintOrderDetailId(code)
	return ord
}

func (od *orddServ) SaveOrderDetail(OrderDetail models.Orderdetail) *models.Orderdetail {

	ord := od.orddRepo.SaveOrderDetail(OrderDetail)
	return ord
}

func (od *orddServ) UpdateOrderDetail(OrderDetail models.Orderdetail) *models.Orderdetail {
	ord := od.orddRepo.UpdateOrderDetail(OrderDetail)
	return ord
}
