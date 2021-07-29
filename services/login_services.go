package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type loginService interface {
	//PrintLogin() []*models.Employee
	GetEmployeEmail(models.Employee) *models.Employee
}

type logServ struct {
	logRepo repository.LoginReposiotry
}


func NewLoginService(logRepo repository.LoginReposiotry) loginService {
	return &logServ{logRepo: logRepo}
}

// func (es *logServ) PrintLogin() []*models.Employee {
// 	log := es.logRepo.PrintLogin()
// 	return log
// }
//func (emp *logServ) LoginEmployee(Login models.Employee) *models.Employee {
	//empId := emp.logRepo.LoginEmployee(Login)
	//return empId

//}

func (es *logServ) LoginEmployee(Login string) *models.Employee {
	emp := es.logRepo.GetEmployeEmail(email)
	return Login
	if emp == nil {
		return nil
	
		if emp.password, err == Login.Password {

	}
}
