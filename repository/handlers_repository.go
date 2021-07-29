package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type LoginReposiotry interface {
	PrintLogin() []*models.Employee
	SaveLogin(models.Employee) *models.Employee
}

func NewLogRepo() LoginReposiotry {
	return &logRepo{}

}

type logRepo struct{}

func (er *logRepo) PrintLogin() []*models.Employee {

	var login []*models.Employee
	result := gormDB.Limit(10).Find(&login)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	return login
}
func (tr *logRepo) SaveLogin(login models.Employee) *models.Employee {
	b := gormDB.First(&login)
	result := gormDB.Save(&b)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("created", login)
	return &login

}
