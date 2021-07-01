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
	RemoveEmployee(models.Employee,string) *models.Employee
	PatchEmployee(models.Employee) *models.Employee

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

func (emp *empServ) SaveEmployee(Employee models.Employee) *models.Employee{
	// var empId int64 = 0
	empId := emp.empRepo.SaveEmployee(Employee)
	return empId
}
func (emp *empServ) EditEmployee(Employee models.Employee) *models.Employee{
	empId := emp.empRepo.EditEmployee(Employee)
	return empId
}

func (emp *empServ) RemoveEmployee(Employee models.Employee , eid string) *models.Employee{
	empId := emp.empRepo.RemoveEmployee(Employee,eid)
	return empId
}

func (emp *empServ) PatchEmployee(Employee models.Employee) *models.Employee{
	empId := emp.empRepo.PatchEmployee(Employee)
	return empId
}

