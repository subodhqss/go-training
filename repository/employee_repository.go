package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type EmployeeRepository interface {
	PrintEmployee() []*models.Employee
	SaveEmployee(models.Employee) *models.Employee
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
func (tr *employeeRepo) SaveEmployee(employee models.Employee) *models.Employee {

	result := gormDB.Create(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in Save all records")
	}

	fmt.Println(&employee)
	return &employee

}
