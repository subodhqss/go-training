// package service

// import (
// 	"github.com/subodhqss/go-training/models"
// 	"github.com/subodhqss/go-training/repository"
// )

// type loginService interface {
// 	//PrintLogin() []*models.Employee
// 	LoginEmployee(models.Login) *models.Login
// }

// type logServ struct {
// 	logRepo repository.LoginReposiotry
// }

// func NewLoginService(logRepo repository.LoginReposiotry) loginService {
// 	return &logServ{logRepo: logRepo}
// }

// // func (es *logServ) PrintLogin() []*models.Employee {
// // 	log := es.logRepo.PrintLogin()
// // 	return log
// // }
// //func (emp *logServ) LoginEmployee(Login models.Employee) *models.Employee {
// 	//empId := emp.logRepo.LoginEmployee(Login)
// 	//return empId

// //}

// func (es *logServ) LoginEmployee(login models.Login) *models.Login {
// 	emp := es.logRepo.GetEmployeByEmail(login.email)

// 	//return emp

// 	if emp.password == Login.Password {
// 			password, _ := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))

// 	//	if emp.password, err == Login.Password {
// 	//	password, _ := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password)

// 	}
// }

package service

import (
	"log"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"golang.org/x/crypto/bcrypt"
)

type loginService interface {
	LoginEmployee(login *models.Login) *models.Login
}

type logServ struct {
	logRepo repository.LoginReposiotry
}

func NewLoginService(logRepo repository.LoginReposiotry) loginService {
	return &logServ{logRepo: logRepo}
}
func (ls *logServ) LoginEmployee(login *models.Login) *models.Login {
	emp := ls.logRepo.GetEmployeEmail(login.Email)
	if emp == nil {
		return nil
	}
	// compare the password here using bcrypt
	hasPass, err := bcrypt.GenerateFromPassword([]byte(login.Password), 5)
	log.Print("hased Pass :", string(hasPass), emp.Password)
	if err != nil {
		return nil
	}

	log.Print("hased Pass :", string(hasPass), emp.Password)

	if bcrypt.CompareHashAndPassword([]byte(emp.Password), hasPass) != nil {
		return nil
	}
	loginData := &models.Login{
		Email:     emp.Email,
		FirstName: emp.FirstName,
		LastName: emp.LastName,
		Password: emp.Password,

	}
	return loginData
}

