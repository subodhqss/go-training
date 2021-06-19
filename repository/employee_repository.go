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
		Add :             "Meerut",
		Designation :     "Software trainee",
		Company_name :    "Qss Technosoft",
		Date_of_joining : "14-06-2021",
		EmpCode :        "522063",
	}
	return emp
}
