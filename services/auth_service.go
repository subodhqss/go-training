package service

import (
	"log"
	"github.com/dgrijalva/jwt-go"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
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
	err := bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(login.Password))
	if err != nil {
		log.Printf("Error: comparing pass %s and %s", emp.Password, login.Password)
		return nil
	}

	//Creating Token
	var jwtKey = []byte("secret_key")

	claims := &models.Claims{
		Email:          emp.Email,
		FirstName:      emp.FirstName,
		LastName:       emp.LastName,
		EmployeeNumber: emp.EmployeeNumber,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	loginData := &models.Login{
		Email:          emp.Email,
		FirstName:      emp.FirstName,
		LastName:       emp.LastName,
		EmployeeNumber: emp.EmployeeNumber,
		Token:          tokenString,
	}
	return loginData
}
