package repository

import (
	"github.com/subodhqss/go-training/models"
)

type employeReposiotry interface {
	PrintEmploye() *models.Employee
}

func NewEmpRepo() employeReposiotry {
	return &empRepo{}
}

type empRepo struct{}

func (er *empRepo) PrintEmploye() *models.Employee {
	emp := &models.Employee{
		Name :            "ALAMGEER",
		add :             "Meerut",
		designation :     "Software trainee",
		company_name :    "Qss Technosoft",
		date_of_joining : "14-06-2021",
		emp_code :        "522063",
	}
	return emp
}
