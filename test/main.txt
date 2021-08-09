package main

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	emp := getEmployeeByEmailData("sk@gmail.com")

	loginPass := "pass@12"

	err := bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(loginPass))
	if err != nil {
		log.Printf("Error: comparing pass %s and %s", emp.Password, loginPass)
	} else {
		fmt.Println("Pass match")
	}

}

func getEmployeeByEmailData(email string) *models.Employee {
	pass := "$2a$05$lcjlRvKB5XsDt3Alkkpbd.m2YD4ZeIs.n7BfB1xTC8r.KjvtYYf9e"
	emp := &models.Employee{
		Password: pass,
	}
	return emp

}
