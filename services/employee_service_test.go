package service

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
	// b, _ := json.MarshalIndent(res, " ", " ")
	// fmt.Println(string(b))
	assert.NotNil(t, res)
	assert.Equal(t, emp, res)
}
