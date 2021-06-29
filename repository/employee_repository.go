package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type EmployeReposiotry interface {
	PrintEmploye() []*models.Employee
	SaveEmployee(models.Employee) *models.Employee
	EditEmployee(models.Employee) *models.Employee
	RemoveEmployee(models.Employee) *models.Employee

}

func NewEmpRepo() EmployeReposiotry {
	return &empRepo{}

}

type empRepo struct{}

func (er *empRepo) PrintEmploye() []*models.Employee {

	var employee []*models.Employee
	result := gormDB.Limit(10).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	return employee
}

func (tr *empRepo) SaveEmployee(employee models.Employee) *models.Employee {
	b := gormDB.First(&employee)
	result := gormDB.Save(&b)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println(&employee)
	return &employee

}

func(tr *empRepo) EditEmployee(employee models.Employee) *models.Employee{

	result := gormDB.Model(&employee).Where("employeeNumber", employee.EmployeeNumber).Updates(employee)
    if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(&employee)
	return &employee

}
func(tr *empRepo) RemoveEmployee(employee models.Employee) *models.Employee{

	result := gormDB.Model(&employee).Where("employeeNumber", 75).Delete(employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(&employee)
	return &employee
}