package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var employeService = service.NewEmployeService(repository.NewEmpRepo())

func GetEmploye(rw http.ResponseWriter, r *http.Request) {
	data := employeService.PrintEmploye()

	jsonData, _ := json.Marshal(data)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreateEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.SaveEmployee(Employee)
}

func UpdateEmployee(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.EditEmployee(Employee)	
}

func DeleteEmployee(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	vars := mux.Vars(r)["eid"]
	employeService.RemoveEmployee(Employee,vars)
	json.NewEncoder(rw).Encode("Employee is deleted!")

}

func PatchEmployee(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.PatchEmployee(Employee)
}
