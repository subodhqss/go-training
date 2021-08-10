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
		EmployeeNumber: 111,
		LastName:       "Diane",
		FirstName:      "Murphy",
		Extension:      "x5800",
		Email:          "dmurphy@classicmodelcars.com",
		OfficeCode   : "1" ,
		ReportsToId: 1002,
		JobTitle:    "President",
	}


	res := repo.SaveEmployee(*employeeDummy)
	assert.Equal(t,res,employeeDummy)

	errCheck := repo.SaveEmployee(models.Employee{})
	assert.NotNil(t,errCheck)

}