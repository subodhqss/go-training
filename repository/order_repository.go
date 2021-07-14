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
	UpdateO(models.Order) *models.Order
}

func NewOrdRepo() OrderReposiotry {

	return &ordRepo{}

}

type ordRepo struct{}

//Order model functions
func (oe *ordRepo) PrintOrder() []*models.Order {
	var order []*models.Order
	result := gormDB.Preload("CustomerDetail").Preload("OrderDetails").Limit(10).Find(&order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return order
}
func (oe *ordRepo) PrintOrderId(code string) *models.Order {
	o_id, _ := strconv.ParseInt(code, 0, 64) //type conversion
	var order *models.Order

	result := gormDB.Preload("CustomerDetail").Preload("OrderDetails").Where("orderNumber", o_id).Find(&order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("No error")
	return order

}

func (oe *ordRepo) SaveOrder(order models.Order) *models.Order {

	result := gormDB.Create(&order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", order)
	return &order
}

func (oe *ordRepo) UpdateOrder(order models.Order) *models.Order {
	result := gormDB.Model(&order).Where("orderNumber", order.OrderNumber).Updates(order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", order)
	return &order
}

func (oe *ordRepo) UpdateO(order models.Order) *models.Order {
	result := gormDB.Model(&order).Where("orderNumber", order.OrderNumber).Updates(order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", order)
	return &order
}
