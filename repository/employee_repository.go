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
	SaveEmployee(models.Employee) *models.Employee
	EditEmployee(models.Employee) *models.Employee
	UpdatePatch(models.Employee) *models.Employee
	DeleteEmployee(models.Employee, string) *models.Employee

	
}

func NewEmpRepo() EmployeReposiotry {
	return &empRepo{}

}

type empRepo struct{}

func (er *empRepo) PrintEmploye() []*models.Employee {

	var employee []*models.Employee
	result := gormDB.Preload("OfficeDetail").Preload("ReportsTo").Limit(10).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	return employee
}

func (er *empRepo) PrintEmployeId(eid string) *models.Employee {
	var employee *models.Employee
	e_id, _ := strconv.ParseInt(eid, 0, 64)

	result := gormDB.Preload("OfficeDetail").Preload("ReportsTo").Where("employeeNumber", e_id).Find(&employee)
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

	fmt.Println("created",employee)
	return &employee

}

func (tr *empRepo) EditEmployee(employee models.Employee) *models.Employee {

	result := gormDB.Model(&employee).Where("employeeNumber", employee.EmployeeNumber).Updates(employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("updated>>>>",employee)
	return &employee
}

func (tr *empRepo) DeleteEmployee(employee models.Employee, eid string) *models.Employee {

	e_id, _ := strconv.ParseInt(eid, 0, 64) //type conversion

	result := gormDB.Where("employeeNumber", e_id).Delete(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("Deleted  !")
	return &employee
}

func (tr *empRepo) UpdatePatch(employee models.Employee) *models.Employee {

	result := gormDB.Model(&employee).Where("employeeNumber", employee.EmployeeNumber).Updates(employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	fmt.Println("update Patch",employee)
	return &employee
}
