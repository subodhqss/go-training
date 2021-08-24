package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/go-training/mocks"
	"github.com/subodhqss/go-training/models"
)

func TestPrintEmployeeById(t *testing.T) {
	repo := new(mocks.EmployeReposiotry)
	empServ := NewEmployeService(repo)
	eid := "4"
	emp := &models.Employee{
		EmployeeNumber: 4,
		FirstName:      "Alam",
		LastName:       "Alam",
	}
	repo.On("PrintEmployeId", eid).Return(emp)

	res := empServ.PrintEmployeId(eid)
	assert.NotNil(t, res)
	assert.Equal(t, emp, res)
}

func TestPrintEmployee(t *testing.T) {
	repo := new(mocks.EmployeReposiotry)
	empServ := NewEmployeService(repo)
	emp := []*models.Employee{
		{EmployeeNumber: 4},
	
	}
	repo.On("PrintEmploye").Return(emp)

	res := empServ.PrintEmploye()
	assert.NotNil(t, res)
	assert.Equal(t, emp, res)
}

func TestSaveEmployee(t *testing.T) {
	repo := new(mocks.EmployeReposiotry)
	empServ := NewEmployeService(repo)
	emp :=models.Employee{
		EmployeeNumber: 4,
		LastName:       "Alam",
		FirstName:      "Alam",
		Extension:      "",
		Email:          "alam@xyz.com",
	}
	
	repo.On("SaveEmployee",emp).Return(&emp)
	res:= empServ.SaveEmployee(emp)
	assert.NotNil(t, res)
	assert.Equal(t, &emp, res)
}

func TestUpdateEmployee(t *testing.T) {
	repo := new(mocks.EmployeReposiotry)
	empServ := NewEmployeService(repo)
	emp :=models.Employee{
		EmployeeNumber: 4,
		LastName:       "Alam",
		FirstName:      "Alam",
		Extension:      "",
		Email:          "alam@xyz.com",
	}
	
	repo.On("UpdateEmployee",emp).Return(&emp)
	res:= empServ.UpdateEmployee(emp)
	assert.NotNil(t, res)
	assert.Equal(t, &emp, res)
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.EmployeReposiotry)
	empServ := NewEmployeService(repo)
	emp :=models.Employee{
		EmployeeNumber: 4,
		LastName:       "Alam",
		FirstName:      "Alam",
		Extension:      "",
		Email:          "alam@xyz.com",
	}
	
	repo.On("Update",emp).Return(&emp)
	res:= empServ.Update(emp)
	assert.NotNil(t, res)
	assert.Equal(t, &emp, res)
}
func TestDeleteEmployee(t *testing.T) {
	repo := new(mocks.EmployeReposiotry)
	empServ := NewEmployeService(repo)
	eid:="4"
	emp :=models.Employee{
		EmployeeNumber: 4,
		LastName:       "Alam",
		FirstName:      "Alam",
		Extension:      "",
		Email:          "alam@xyz.com",
	}
	
	repo.On("DeleteEmployee",emp,eid).Return(&emp)
	res:= empServ.DeleteEmployee(emp,eid)
	assert.NotNil(t, res)
	assert.Equal(t, &emp, res)
}