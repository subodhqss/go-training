package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Employee struct {
	Employeenumber string `json:"employeenumber"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Officecode     string `json:"officecode"`
}

func main() {
	csvFile, err := os.Open("emp.csv")
	if err != nil {
		fmt.Println(err)
	}
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		emp := Employee{
			Employeenumber: line[0],
			Firstname:      line[1],
			Lastname:       line[2],
			Email:          line[3],
			Officecode:     line[4],
		}
		fmt.Println(emp.Employeenumber + " " + emp.Firstname + " " + emp.Lastname + " " + emp.Email + " " + emp.Officecode)
	}
}
