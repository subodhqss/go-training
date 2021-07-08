package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type PaymentService interface {
	PrintPayment() []*models.Payment
	PrintPaymentId(string) []*models.Payment
	SavePayment(models.Payment) *models.Payment
	UpdatePayment(models.Payment) *models.Payment
}

type payServ struct {
	payRepo repository.PaymentReposiotry
}

func NewPaymentService(payRepo repository.PaymentReposiotry) PaymentService {
	return &payServ{payRepo: payRepo}
}


func (pys *payServ) PrintPayment() []*models.Payment {
	pay := pys.payRepo.PrintPayment()
	return pay
}

func (pys *payServ) PrintPaymentId(code string) []*models.Payment {
	pay := pys.payRepo.PrintPaymentId(code)
	return pay
}

func (pys *payServ) SavePayment(Payment models.Payment) *models.Payment {

	pay := pys.payRepo.SavePayment(Payment)
	return pay
}

func (pys *payServ) UpdatePayment(Payment models.Payment) *models.Payment {
	pay := pys.payRepo.UpdatePayment(Payment)
	return pay
}
