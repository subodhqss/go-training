package repository

import (
	"github.com/subodhqss/go-training/models"
)

type EmployeReposiotry interface {
	PrintEmploye() *models.Employee
}

func NewEmpRepo() EmployeReposiotry {
	return &empRepo{}
}

type empRepo struct{}

func (er *empRepo) PrintEmploye() *models.Employee {
	emp := &models.Employee{
		Name:            "Gurneet Kaur",
		Add:             "Meerut",
		Designation:     "Software trainee",
		Company_name:    "Qss Technosoft",
		Date_of_joining: "14-06-2021",
		EmpCode:         "522064",
	}
	return emp
}
