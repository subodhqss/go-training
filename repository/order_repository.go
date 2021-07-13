package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/subodhqss/go-training/models"
)

type OrderReposiotry interface {
	PrintOrder() []*models.Order
	PrintOrderId(string) *models.Order
	SaveOrder(models.Order) *models.Order
	UpdateOrder(models.Order) *models.Order
}

func NewOrdRepo() OrderReposiotry {
	return &ordRepo{}
}

type ordRepo struct{}

func (er *ordRepo) PrintOrder() []*models.Order {
	var Order []*models.Order
	result := gormDB.Limit(10).Find(&Order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(result)
	return Order
}

func (er *ordRepo) PrintOrderId(code string) *models.Order {

	o_id, _ := strconv.ParseInt(code, 0, 64)
	Order := &models.Order{}
	result := gormDB.Preload("CustomerDetail").Preload("OrderDetails").Where("orderNumber", o_id).Find(Order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error ", result, code)
	return Order
}

func (tr *ordRepo) SaveOrder(Order models.Order) *models.Order {

	result := gormDB.Create(&Order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created ", Order)
	return &Order
}

func (tr *ordRepo) UpdateOrder(Order models.Order) *models.Order {
	result := gormDB.Model(&Order).Where("orderNumber", Order.OrderNumber).Updates(Order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated", Order)
	return &Order
}
