package repository

import (
	
	"fmt"
	"log"
	"strconv"

	"github.com/subodhqss/go-training/models"
)

type OrderDetailReposiotry interface {
	PrintOrderDetail() []*models.Orderdetail
	PrintOrderDetailId(string) *models.Orderdetail
	SaveOrderDetail(models.Orderdetail) *models.Orderdetail
	UpdateOrderDetail(models.Orderdetail) *models.Orderdetail

}

func NewOrddRepo() OrderDetailReposiotry {
	return &orddRepo{}
}

type orddRepo struct{}

func (er *orddRepo) PrintOrderDetail() []*models.Orderdetail {
	var OrderDetail []*models.Orderdetail
	result := gormDB.Limit(10).Find(&OrderDetail)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(result)
	return OrderDetail
}

func (er *orddRepo) PrintOrderDetailId(code string) *models.Orderdetail {

	o_id, _ := strconv.ParseInt(code, 0, 64) //type conversion
	OrderDetail := &models.Orderdetail{} 
	result := gormDB.Where("OrderNumber", o_id).Find(OrderDetail)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method", result, code)
	return OrderDetail
}

func (tr *orddRepo) SaveOrderDetail(OrderDetail models.Orderdetail) *models.Orderdetail {

	result := gormDB.Create(&OrderDetail)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", OrderDetail)
	return &OrderDetail
}

func (tr *orddRepo) UpdateOrderDetail(OrderDetail models.Orderdetail) *models.Orderdetail {
	result := gormDB.Model(&OrderDetail).Where("orderNumber", OrderDetail.OrderNumber).Updates(OrderDetail)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", OrderDetail)
	return &OrderDetail
}

