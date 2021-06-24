package repository

import (
	"fmt"

	"github.com/subodhqss/go-training/models"
)

type EmployeReposiotry interface {
	PrintEmploye() *models.Employee
}

func NewEmpRepo() EmployeReposiotry {
	return &empRepo{}
}

type empRepo struct{}

// var employee models.Employee
// var employees []models.Employee

func (er *empRepo) PrintEmploye() *models.Employee {
	// var employee models.Employee
	// var employees [] empRepo

	// result := map[string]interface{}{}
	// gormDB.Table("employees").Take(&result)

	// var Employee [] models.Employee

	// db.Find(&employee)

	e := &models.Employee{}

	res := gormDB.Limit(10).Take(e)

	fmt.Println("res ", res.RowsAffected)

	// gormDB.Raw("select * from employees limit 10").Scan(&e)

	// gormDB.Select("Empnumber", "LastName", "FirstName",
	// "Extension", "Email", "OfficeCode", "ReportTo", "JobTitle").Create(&er)

	// emp := &models.Employee{
	// 	Name:            "ALAMGEER",
	// 	Add:             "Meerut",
	// 	Designation:     "Software trainee",
	// 	Company_name:    "Qss Technosoft",
	// 	Date_of_Joining: "14-06-2021",
	// 	EmpCode:         "522063",
	// }
	return e
}
