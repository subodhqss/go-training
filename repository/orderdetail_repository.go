package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type OrderdetailReposiotry interface {
	PrintOrderdetail() []*models.Orderdetail
	SaveDetail(models.Orderdetail) *models.Orderdetail
	UpdateDetail(models.Orderdetail) *models.Orderdetail
	UpdateD(models.Orderdetail) *models.Orderdetail
}

func NewOrderRepo() OrderdetailReposiotry {

	return &orderRepo{}

}

type orderRepo struct{}

//orderdetail functions
func (or *orderRepo) PrintOrderdetail() []*models.Orderdetail {
	var orderdetail []*models.Orderdetail
	result := gormDB.Limit(10).Find(&orderdetail)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return orderdetail
}

/*func (cr *custRepo) PrintCustomerNumber(number int) *models.Customer {
	var customer *models.Customer
	e_id, _ := strconv.ParseInt(eid, 0, 64) //type conversion
	result := gormDB.Preload("ReportsTo").Where("employeeNumber", e_id).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method")
	return customer
}*/

func (or *orderRepo) SaveDetail(orderdetail models.Orderdetail) *models.Orderdetail {

	result := gormDB.Create(&orderdetail)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", orderdetail)
	return &orderdetail
}

func (or *orderRepo) UpdateDetail(orderdetail models.Orderdetail) *models.Orderdetail {
	result := gormDB.Model(&orderdetail).Where("orderNumber", orderdetail.OrderNumber).Updates(orderdetail)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", orderdetail)
	return &orderdetail
}

func (or *orderRepo) UpdateD(orderdetail models.Orderdetail) *models.Orderdetail {
	result := gormDB.Model(&orderdetail).Where("orderNumber", orderdetail.OrderNumber).Updates(orderdetail)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", orderdetail)
	return &orderdetail
}