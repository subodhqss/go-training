package repository

import (
	// "log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/go-training/models"
)

func TestSaveEmployee1(t *testing.T){
	repo := NewEmpRepo()
	employeeDummy := &models.Employee{
		
		LastName:       "Diane",
		FirstName:      "Murphy",
		Extension:      "x5800",
		Email:          "dmurphy@classicmodelcars.com",
		OfficeCode   : "1" ,
		ReportsToId: 1002,
		JobTitle:    "President",
	}


	res := repo.SaveEmployee(*employeeDummy)
	employeeDummy.EmployeeNumber = res.EmployeeNumber
	assert.Equal(t,employeeDummy,res)

	employeeDummy1:=models.Employee{
		EmployeeNumber: 999,
	}

	errCheck := repo.SaveEmployee(employeeDummy1)
	assert.Nil(t,errCheck)


}