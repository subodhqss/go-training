package services

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type PaymentService interface {
	PrintPayment() []*models.Payment
	 SavePayment(models.Payment) *models.Payment
	// UpdatePayment(models.Payment) *models.Payment
	// UpdatePa(models.Payment) *models.Payment
}

type payServ struct {
	payRepo repository.PaymentReposiotry
}

func NewPaymentService(payRepo repository.PaymentReposiotry) PaymentService {
	return &payServ{payRepo: payRepo}
}


func (ps *payServ) PrintPayment() []*models.Payment {
	pay := ps.payRepo.PrintPayment()
	return pay
}

func (pay *payServ) SavePayment(Payment models.Payment) *models.Payment {

	chkNo := pay.payRepo.SavePayment(Payment)
	return chkNo
}
