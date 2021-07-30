package controllers

import (
	"encoding/json"
	"log"

	//"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	service "github.com/subodhqss/go-training/services"
	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/crypto/bcrypt"
)

var employeService = service.NewEmployeService(repository.NewEmpRepo())

func GetEmploye(rw http.ResponseWriter, r *http.Request) {
	data := employeService.PrintEmploye()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func GetEmployeId(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["eid"]
	data := employeService.PrintEmployeId(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreateEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	hasPass, err := bcrypt.GenerateFromPassword([]byte(Employee.Password), 5)
	log.Print("hased Pass :", string(hasPass), Employee.Password)
	if err != nil {
		return
	}
	Employee.Password = string(hasPass)
	data := employeService.SaveEmployee(Employee)
	json.NewEncoder(rw).Encode(data)

}

// func CreateEmployee(rw http.ResponseWriter, r *http.Request) {
// 	var Employee models.Employee
// 	err := json.NewDecoder(r.Body).Decode(Employee)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Employee.Password), 8)
// 	if _, err = gormDB.Query("insert", Employee.Email, string(hashedPassword)); err != nil {
// 		rw.WriteHeader(http.StatusInternalServerError)
// 	}
// }

func UpdateEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.EditEmployee(Employee)
	json.NewEncoder(rw).Encode("Updated....")
	json.NewEncoder(rw).Encode(Employee)

}

func DeleteEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	vars := mux.Vars(r)["eid"]
	employeService.DeleteEmployee(Employee, vars)

}

func UpdatePatch(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.UpdatePatch(Employee)

}
