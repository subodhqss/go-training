package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type employeService interface {
	PrintEmploye() *models.Employee
}

type empSrv struct {
	empRepo repository.employeReposiotry
}

func NewEmployeService(empRepo repository.employeReposiotry) employeService {
	return &empServ{empRepo: empRepo}
}

func (es *empSrv) PrintEmploye() *models.Employee {
	emp := es.empRepo.PrintEmploye()
	return emp
}
