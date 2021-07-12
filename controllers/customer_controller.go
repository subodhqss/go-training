package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var customerService = service.CustomerService(repository.NewCustRepo())

func GetCustomer(rw http.ResponseWriter, r *http.Request) {
	data := customerService.PrintCustomer()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}


/*func GetCustomerNumber(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["cno"]
	data := customerService.PrintCustomerNumber(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}*/
func CreateCustomer(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Customer models.Customer
	json.NewDecoder(r.Body).Decode(&Customer)
	customerService.SaveCustomer(Customer)
	json.NewEncoder(rw).Encode("Customer is Created Succesfully !")
	json.NewEncoder(rw).Encode(Customer)
}

func UpdateCustomer(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Customer models.Customer
	json.NewDecoder(r.Body).Decode(&Customer)
	customerService.UpdateCustomer(Customer)
	json.NewEncoder(rw).Encode("Customer is updated by PUT method !")
	json.NewEncoder(rw).Encode(Customer)
}

func UpdateC(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Customer models.Customer
	json.NewDecoder(r.Body).Decode(&Customer)
	customerService.UpdateC(Customer)
	json.NewEncoder(rw).Encode("Customer is updated by PATCH method !")
	json.NewEncoder(rw).Encode(Customer)
}
