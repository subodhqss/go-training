package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/subodhqss/go-training/models"
)

type EmployeReposiotry interface {
	PrintEmploye() []*models.Employee
	PrintEmployeId(string) *models.Employee
	//PrintEmployeEmail(string) *models.Employee
	SaveEmployee(models.Employee) *models.Employee
	UpdateEmployee(models.Employee) *models.Employee
	Update(models.Employee) *models.Employee
	DeleteEmployee(models.Employee, string) *models.Employee
	SaveEmployee1(models.Employee, string) *models.Employee
}

func NewEmpRepo() EmployeReposiotry {
	return &empRepo{}
	
}

type empRepo struct{}

//Employee model functions
func (er *empRepo) PrintEmploye() []*models.Employee {
	var employee []*models.Employee
	result := gormDB.Preload("ReportsTo").Limit(10).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return employee
}

func (er *empRepo) PrintEmployeId(eid string) *models.Employee {
	var employee *models.Employee
	e_id, _ := strconv.ParseInt(eid, 0, 64) //type conversion

	result := gormDB.Preload("OfficeDetail").Preload("ReportsTo").Where("employeeNumber", e_id).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method(id)")
	return employee
}
func (tr *empRepo) SaveEmployee(employee models.Employee) *models.Employee {

	result := gormDB.Create(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Created new entry succesfull !", employee)
	return &employee
}

func (tr *empRepo) UpdateEmployee(employee models.Employee) *models.Employee {
	result := gormDB.Model(&employee).Where("employeeNumber", employee.EmployeeNumber).Updates(employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", employee)
	return &employee
}

func (tr *empRepo) Update(employee models.Employee) *models.Employee {
	result := gormDB.Model(&employee).Where("employeeNumber", employee.EmployeeNumber).Updates(employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("Updated Succesfull !", employee)
	return &employee
}

func (tr *empRepo) DeleteEmployee(employee models.Employee, eid string) *models.Employee {

	e_id, _ := strconv.ParseInt(eid, 0, 64) //type conversion

	result := gormDB.Where("employeeNumber", e_id).Delete(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Employee number", e_id, "Deleted Succesfully !")
	return &employee
}
func (tr *empRepo) SaveEmployee1(employee models.Employee) *models.Employee {

	result := &models.Employee{
		EmployeeNumber: 109,
		LastName:      "kaur"
		FirstName:     "Neet"
		Extension:     "abcc"
		Email:         "neet@gmail.com"
		ReportsToId:   112,
		JobTitle:      "trainee"
	}
	
}
