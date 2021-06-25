package controllers

import (
	"encoding/json"
	"net/http"

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
