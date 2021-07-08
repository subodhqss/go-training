package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type PaymentReposiotry interface {
	PrintPayment() []*models.Payment
	SavePayment(models.Payment) *models.Payment
	UpdatePayment(models.Payment) *models.Payment
	UpdatePa(models.Payment) *models.Payment

}
func NewPayRepo() PaymentReposiotry {
	
	return &payRepo{}

}

type payRepo struct{}

//Payment model functions
func (pr *payRepo) PrintPayment() []*models.Payment {
	var payment []*models.Payment
	result := gormDB.Limit(10).Find(&payment)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return payment
}


func (pr *payRepo) SavePayment(payment models.Payment) *models.Payment {

	result := gormDB.Create(&payment)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", payment)
	return &payment
}

func (pr *payRepo) UpdatePayment(payment models.Payment) *models.Payment {
	result := gormDB.Model(&payment).Where("checkNumber", payment.CheckNumber).Updates(payment)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", payment)
	return &payment
}

func (pr *payRepo) UpdatePa(payment models.Payment) *models.Payment {
	result := gormDB.Model(&payment).Where("checkNumber", payment.CheckNumber).Updates(payment)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", payment)
	return &payment
}







