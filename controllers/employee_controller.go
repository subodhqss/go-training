package controllers

import (
	"encoding/json"
	// "log"
	"net/http"

	 "github.com/gorilla/mux"
	 "github.com/subodhqss/go-training/models"
	 "github.com/subodhqss/go-training/repository"
	 "github.com/subodhqss/go-training/services"
	//   "golang.org/x/crypto/bcrypt"
)

var employeService = services.NewEmployeService(repository.NewEmpRepo())

// func CreateEmployee(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Set("Content-Type", "application/json")
// 	var Employee models.Employee
// 	json.NewDecoder(r.Body).Decode(&Employee)

// 	hasPass, err := bcrypt.GenerateFromPassword([]byte(Employee.Password), 5)
// 	log.Print("hased Pass :", string(hasPass), Employee.Password)
// 	if err != nil {
// 		return
// 	}
// 	Employee.Password = string(hasPass)
// 	data := employeService.SaveEmployee(Employee)
// 	json.NewEncoder(rw).Encode(data)
// 	// log.Print("hased Pass :", string(hasPass), emp.Password)

// }

func GetEmploye(rw http.ResponseWriter, r *http.Request) {
	data := employeService.PrintEmploye()
	jsonData, _ := json.Marshal(data)                                  
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}



func UpdateEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.UpdateEmployee(Employee)
	json.NewEncoder(rw).Encode("Employee is updated by PUT method !")
	json.NewEncoder(rw).Encode(Employee)
}

func GetEmployeId(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["employeeNumber"]
	data := employeService.PrintEmployeId(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func DeleteEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	vars := mux.Vars(r)["employeeNumber"]
	employeService.DeleteEmployee(Employee, vars)
	json.NewEncoder(rw).Encode("Employee number: " + vars + " Deleted Succesfully !")
}
