package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type PaymentService interface {
	PrintPayment() []*models.Payment
	PrintPaymentId(string) *models.Payment
	SavePayment(models.Payment) *models.Payment
	UpdatePayment(models.Payment) *models.Payment
}

type payServ struct {
	payRepo repository.PaymentReposiotry
}

func NewPaymentService(payRepo repository.PaymentReposiotry) PaymentService {
	return &payServ{payRepo: payRepo}
}

func (od *payServ) PrintPayment() []*models.Payment {
	ordd := od.payRepo.PrintPayment()
	return ordd
}

func (od *payServ) PrintPaymentId(code string) *models.Payment {
	ord := od.payRepo.PrintPaymentId(code)
	return ord
}

func (od *payServ) SavePayment(Payment models.Payment) *models.Payment {

	ord := od.payRepo.SavePayment(Payment)
	return ord
}

func (od *payServ) UpdatePayment(Payment models.Payment) *models.Payment {
	ord := od.payRepo.UpdatePayment(Payment)
	return ord
}
