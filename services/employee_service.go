package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type employeService interface {
	PrintEmploye() []*models.Employee
	PrintEmployeId(string) *models.Employee
	SaveEmployee(models.Employee) *models.Employee
	EditEmployee(models.Employee) *models.Employee
	UpdatePatch(models.Employee) *models.Employee
	DeleteEmployee(models.Employee, string) *models.Employee
}

type empServ struct {
	empRepo repository.EmployeReposiotry
}

func NewEmployeService(empRepo repository.EmployeReposiotry) employeService {
	return &empServ{empRepo: empRepo}
}

func (es *empServ) PrintEmploye() []*models.Employee {
	emp := es.empRepo.PrintEmploye()
	return emp
}
func (es *empServ) PrintEmployeId(eid string) *models.Employee {
	emp := es.empRepo.PrintEmployeId(eid)
	return emp
}

func (emp *empServ) SaveEmployee(Employee models.Employee) *models.Employee {
	empId := emp.empRepo.SaveEmployee(Employee)
	return empId
}
func (emp *empServ) EditEmployee(Employee models.Employee) *models.Employee {
	empId := emp.empRepo.EditEmployee(Employee)
	return empId
}

func (em *empServ) DeleteEmployee(Employee models.Employee, eid string) *models.Employee {
	empId := em.empRepo.DeleteEmployee(Employee, eid)
	return empId
}
func (emp *empServ) UpdatePatch(Employee models.Employee) *models.Employee {
	empId := emp.empRepo.UpdatePatch(Employee)
	return empId
}

