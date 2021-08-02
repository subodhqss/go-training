package service

import (
	"log"
	"fmt"
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
	//password := "pswd1"
	err := bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(login.Password))
	if err != nil {
		log.Printf("Error: comparing pass %s and %s", emp.Password, login.Password)
	} else {
		fmt.Println("Pass match")
	}
	/*hasPass, err := bcrypt.GenerateFromPassword([]byte(login.Password), 5)
	log.Print("hased Pass :", string(hasPass), emp.Password)
	if err != nil {
		return nil
	}

	log.Print("hased Pass :", string(hasPass), emp.Password)

	if bcrypt.CompareHashAndPassword([]byte(emp.Password), hasPass) != nil {
		return nil
	}I*/
	loginData := &models.Login{
		Email:     emp.Email,
		FirstName: emp.FirstName,
		LastName:  emp.LastName,
		EmployeeNumber: emp.EmployeeNumber,
		
	}
	return loginData
}
