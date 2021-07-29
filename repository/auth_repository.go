package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type LoginReposiotry interface {
	//PrintLogin() []*models.Employee
	//SaveLogin(models.Employee) *models.Employee
}

func NewLogRepo() LoginReposiotry {
	return &logRepo{}

}

type logRepo struct{}

func (lr *logRepo) GetEmployeEmail(email string) *models.Employee {
	var employee *models.Employee

	fmt.Println(email) 

	result := gormDB.Preload("OfficeDetail").Preload("ReportsTo").Where("email", email).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method")
    return employee
}

