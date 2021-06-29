package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type employeService interface {
	PrintEmploye() []*models.Employee
	SaveEmployee(models.Employee) *models.Employee
	EditEmployee(models.Employee) *models.Employee
	DeleteEmployee(models.Employee, string)*models.Employee

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

func (emp *empServ) SaveEmployee(Employee models.Employee) *models.Employee{
	// var empId int64 = 0
	empId := emp.empRepo.SaveEmployee(Employee)
	return empId
}
func (emp *empServ) EditEmployee(Employee models.Employee) *models.Employee{
	empId := emp.empRepo.EditEmployee(Employee)
	return empId
}

func(em *empServ) DeleteEmployee(Employee models.Employee, eid string) *models.Employee{

	
	empId := em.empRepo.DeleteEmployee(Employee, eid)
	return empId
}