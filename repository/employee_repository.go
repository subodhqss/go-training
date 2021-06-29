package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type EmployeReposiotry interface {
	PrintEmploye() []*models.Employee
	SaveEmployee(models.Employee) *models.Employee
	UpdateEmployee(models.Employee) *models.Employee
	DeleteEmployee(models.Employee) *models.Employee

}

func NewEmpRepo() EmployeReposiotry {
	return &empRepo{}
	
}

type empRepo struct{}



func (er *empRepo) PrintEmploye()[] *models.Employee {
	

	var employee []*models.Employee
	result := gormDB.Limit(10).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	
	return employee
}

func (tr *empRepo) SaveEmployee(employee models.Employee) *models.Employee {
	

	result:= gormDB.Create(&employee)
	if err := result.Error; err != nil {
			log.Print("Error in getting all records")
		}

	fmt.Println("Created new entry succesfull !",employee)
	return &employee
	
}

func (tr *empRepo) UpdateEmployee(employee models.Employee) *models.Employee {
	

	result := gormDB.Model(&employee).Where("employeeNumber", employee.EmployeeNumber).Updates(employee)

	
	
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	
	fmt.Println("Updated Succesfull !",employee)
	return &employee
	
}

func (tr *empRepo) DeleteEmployee(employee models.Employee) *models.Employee{

	eid:=75
	result:= gormDB.Where("employeeNumber",eid).Delete(&employee)
	if err := result.Error; err != nil {
			log.Print("Error in getting all records")
		}

	fmt.Println("Deleted Succesfull")
	return &employee
}