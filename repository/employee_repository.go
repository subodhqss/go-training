package repository

import (
	"log"

	"github.com/subodhqss/go-training/models"
)

type EmployeeRepository interface {
	PrintEmployee() []*models.Employee
	SaveEmployee(models.Employee) int64
}

func NewEmpRepo() EmployeeRepository {
	return &employeeRepo{}
}

type employeeRepo struct{}

func (tr *employeeRepo) PrintEmployee() []*models.Employee {

	var employee []*models.Employee
	result := gormDB.Limit(10).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return employee
}
func (tr *employeeRepo) SaveEmployee(models.Employee) int64 {
	var employee models.Employee
	result := gormDB.Create(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in save all records")
	}
	return int64(employee.EmployeeNumber)
}
