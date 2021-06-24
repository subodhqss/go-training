package repository

import (
	"log"

	"github.com/subodhqss/go-training/models"
)

type EmployeReposiotry interface {
	PrintEmploye() []*models.Employee
}

func NewEmpRepo() EmployeReposiotry {
	return &empRepo{}
	
}

type empRepo struct{}



func (er *empRepo) PrintEmploye()[] *models.Employee {
	

	var employee []*models.Employee
	result := gormDB.Limit(10).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

		return employee
}