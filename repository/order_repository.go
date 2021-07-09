package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type OrderReposiotry interface {
	PrintOrder() []*models.Order
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
	result := gormDB.Limit(10).Find(&order)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
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
