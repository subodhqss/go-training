package repository

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/go-training/models"
)

func TestSaveEmployee1(t *testing.T) {
	repo := NewEmpRepo()
	employeeDummy := &models.Employee{

		LastName:    "Diane",
		FirstName:   "Murphy",
		Extension:   "x5800",
		Email:       "dmurphy@classicmodelcars.com",
		OfficeCode:  "1",
		ReportsToId: 1002,
		JobTitle:    "President",
	}

	res := repo.SaveEmployee(*employeeDummy)
	employeeDummy.EmployeeNumber = res.EmployeeNumber
	assert.Equal(t, employeeDummy, res)

	employeeDummy1 := models.Employee{
		EmployeeNumber: 999,
	}
	errCheck := repo.SaveEmployee(employeeDummy1)
	assert.Nil(t, errCheck)
}

func TestGetEmployee(t *testing.T) {
	repo := NewEmpRepo()

	var dummyEmployee []*models.Employee
	result := gormDB.Preload("OfficeDetail").Preload("ReportsTo").Limit(10).Find(&dummyEmployee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	res := repo.PrintEmploye()
	assert.Equal(t, dummyEmployee, res)

}

func TestGetEmployeeId(t *testing.T) {
	repo := NewEmpRepo()

	var dummyEmployee *models.Employee
	result := gormDB.Preload("OfficeDetail").Preload("ReportsTo").Where("employeeNumber", 99).Find(&dummyEmployee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	res := repo.PrintEmployeId("99")
	assert.Equal(t, dummyEmployee, res)

	errCheck := repo.PrintEmployeId("jsdasj") //its not working
	assert.Nil(t, errCheck)

}

func TestUpdateEmployee(t *testing.T) {
	repo := NewEmpRepo()
	employeeDummy := &models.Employee{
		EmployeeNumber: 99,
		LastName:       "Alam",
		FirstName:      "Murphy",
		Extension:      "x5800",
		Email:          "dmurphy@classicmodelcars.com",
		OfficeCode:     "1",
		ReportsToId:    1002,
		JobTitle:       "President",
	}

	res := repo.UpdateEmployee(*employeeDummy)
	assert.Equal(t, employeeDummy, res)

	errModel := models.Employee{
		EmployeeNumber: 000,
		LastName:       "Alam",
		FirstName:      "Murphy",
		Extension:      "x5800",
		Email:          "dmurphy@classicmodelcars.com",
		OfficeCode:     "1",
		ReportsToId:    1002,
		JobTitle:       "President",
	}

	errCheck := repo.UpdateEmployee(errModel)
	assert.Nil(t,errCheck)//not working

}
