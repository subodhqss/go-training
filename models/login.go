package models

import "github.com/dgrijalva/jwt-go"

type Login struct {
	EmployeeNumber int    `json:"employee_number" `
	LastName       string `json:"last_name"`
	FirstName      string `json:"first_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Token          string `json:"token"`
	jwt.StandardClaims
}
