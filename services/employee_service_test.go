package service

import (
	"encoding/json"
	"fmt"
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
		FirstName: "Gurneet",
		LastName:  "Kaur",
	}
	repo.On("PrintEmployeId", eid).Return(emp)

	res := empServ.PrintEmployeId(eid)
	b, _ := json.MarshalIndent(res, " ", " ")
	fmt.Println(string(b))
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
	emp := models.Employee{
		EmployeeNumber: 4,
		FirstName: "Gurneet",
		LastName: "Kaur",
	}
	repo.On("SaveEmployee",emp).Return(&emp)

	res := empServ.SaveEmployee(emp)

	assert.NotNil(t, res)
	assert.Equal(t, &emp, res)
}
func TestUpdateEmployee(t *testing.T) {
	repo := new(mocks.EmployeReposiotry)
	empServ := NewEmployeService(repo)
	emp := models.Employee{
		EmployeeNumber: 4,
		FirstName: "Gurneet",
		LastName: "Kaur",
	}
	repo.On("UpdateEmployee",emp).Return(&emp)

	res := empServ.UpdateEmployee(emp)

	assert.NotNil(t, res)
	assert.Equal(t, &emp, res)
}
func TestUpdate(t *testing.T) {
	repo := new(mocks.EmployeReposiotry)
	empServ := NewEmployeService(repo)
	emp := models.Employee{
		EmployeeNumber: 4,
		FirstName: "Gurneet",
		LastName: "Kaur",
	}
	repo.On("Update",emp).Return(&emp)

	res := empServ.Update(emp)

	assert.NotNil(t, res)
	assert.Equal(t, &emp, res)
}


 func TestDeleteEmployee(t *testing.T) {
 	repo := new(mocks.EmployeReposiotry)
 	empServ := NewEmployeService(repo)
	eid := "4"
 	emp := models.Employee{
		 
 		EmployeeNumber: 4,
		FirstName: "Gurneet",
		LastName:   "Kaur",
	}
	repo.On("DeleteEmployee",emp,eid).Return(&emp)
	res := empServ.DeleteEmployee(emp,eid)
	
 	assert.NotNil(t, res)
 	assert.Equal(t, &emp, res)
 }




