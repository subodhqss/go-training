package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var employeeService = services.NewEmployeeService(repository.NewEmpRepo())

func GetEmployee(rw http.ResponseWriter, r *http.Request) {
	data := employeeService.PrintEmp()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreateEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Employee models.Employee
	fmt.Println("Controller  :: ", Employee)
	json.NewDecoder(r.Body).Decode(&Employee)
	employeeService.SaveEmployee(Employee)

}
