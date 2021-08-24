package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	service "github.com/subodhqss/go-training/services"
)

var employeService = service.NewEmployeService(repository.NewEmpRepo())

//Employee model functions
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
	employeService.SaveEmployee(Employee)
	json.NewEncoder(rw).Encode("Employee is Created Succesfully !")
	json.NewEncoder(rw).Encode(Employee)
}
func CreateEmployeeBulk(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee []models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.SaveEmployeeBulk(Employee)
	json.NewEncoder(rw).Encode("Employee is Created Succesfully !")
	json.NewEncoder(rw).Encode(Employee)
}

func CreateEmployeeCSV(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.SaveEmployeeCSV(Employee)
	json.NewEncoder(rw).Encode("Employee is Created Succesfully !")
	json.NewEncoder(rw).Encode(Employee)
	
}


func UpdateEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.UpdateEmployee(Employee)
	json.NewEncoder(rw).Encode("Employee is updated by PUT method !")
	json.NewEncoder(rw).Encode(Employee)
}

func Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	json.NewDecoder(r.Body).Decode(&Employee)
	employeService.Update(Employee)
	json.NewEncoder(rw).Encode("Employee is updated by PATCH method !")
	json.NewEncoder(rw).Encode(Employee)
}

func DeleteEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	vars := mux.Vars(r)["eid"]
	employeService.DeleteEmployee(Employee, vars)
	json.NewEncoder(rw).Encode("Employee number: " + vars + " Deleted Succesfully !")
}

func EmployeeImage(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	eId := mux.Vars(r)["eid"]
	r.ParseMultipartForm(10 << 20)

	file, header, _ := r.FormFile("myFile")
	defer file.Close()

	employeService.UpdateImage(eId,file,header,r.Host)
	json.NewEncoder(rw).Encode("Employee number: "+eId+ " Image updated Succesfully !")

}
