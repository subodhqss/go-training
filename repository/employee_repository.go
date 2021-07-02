package repository

import (
	"strconv"
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type EmployeReposiotry interface {
	PrintEmploye() []*models.Employee
	PrintEmployeId(string) *models.Employee
	SaveEmployee(models.Employee) *models.Employee
	EditEmployee(models.Employee) *models.Employee
	RemoveEmployee(models.Employee,string) *models.Employee
	PatchEmployee(models.Employee) *models.Employee


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
func (er *empRepo) PrintEmployeId(eid string) *models.Employee {

	var employee *models.Employee
	e_id,_ := strconv.ParseInt(eid,0,64)
	result:=  gormDB.Preload("ReportTo").Where("employeeNumber",e_id).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	return employee
}

func (tr *empRepo) SaveEmployee(employee models.Employee) *models.Employee {
	
	result := gormDB.Create(&employee)
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
func(tr *empRepo) RemoveEmployee(employee models.Employee, eid string) *models.Employee{
    e_id,_ := strconv.ParseInt(eid,0,64)
	result := gormDB.Model(&employee).Where("employeeNumber", e_id).Delete(employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(&employee)
	return &employee
}

func(tr *empRepo) PatchEmployee(employee models.Employee) *models.Employee{

	result := gormDB.Model(&employee).Where("employeeNumber", employee.EmployeeNumber).Updates(employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println(&employee)
	return &employee
}