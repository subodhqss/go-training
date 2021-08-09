package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/go-training/models"
)

func TestSaveEmployee1(t *testing.T){
	repo := NewEmpRepo()
	expRes := &models.Employee{
		EmployeeNumber: 112,
		LastName:       "aasda",
		FirstName:      "dada",
		Extension:      "adsdaassd",
		Email:          "faf@sd.com",
		// OfficeCode   : 113 ,
		// OfficeDetail: "fgdg",
		// // ReportsTo  : "dfgd",
		ReportsToId: 113,
		JobTitle:    "hfhfhf",
	}

	res := repo.SaveEmployee1(models.Employee{})
	assert.Equal(t,expRes,res)
}