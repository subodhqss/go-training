package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/subodhqss/go-training/models"
)

type PaymentReposiotry interface {
	PrintPayment() []*models.Payment
	PrintPaymentId(string) *models.Payment
	SavePayment(models.Payment) *models.Payment
	UpdatePayment(models.Payment) *models.Payment
}

func NewPayRepo() PaymentReposiotry {
	return &payRepo{}
}

type payRepo struct{}

func (er *payRepo) PrintPayment() []*models.Payment {
	var Payment []*models.Payment
	result := gormDB.Limit(10).Find(&Payment)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(result)
	return Payment
}

func (er *payRepo) PrintPaymentId(code string) *models.Payment {

	o_id, _ := strconv.ParseInt(code, 0, 64) 
	Payment := &models.Payment{}
	result := gormDB.Where("PaymentNumber", o_id).Find(Payment)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error ", result, code)
	return Payment
}

func (tr *payRepo) SavePayment(Payment models.Payment) *models.Payment {

	result := gormDB.Create(&Payment)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created ", Payment)
	return &Payment
}

func (tr *payRepo) UpdatePayment(Payment models.Payment) *models.Payment {
	result := gormDB.Model(&Payment).Where("customerNumber", Payment.CustomerNumber).Updates(Payment)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated", Payment)
	return &Payment
}
