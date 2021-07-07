package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var customerService = service.NewCustomerService(repository.NewCusRepo())
 
func GetCustomers(rw http.ResponseWriter, r *http.Request) {
	data := customerService.PrintCustomer()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func GetCustomer(rw http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)["code"]
	data := customerService.PrintCustomerId(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreatCustomer(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Customer models.Customer
	json.NewDecoder(r.Body).Decode(&Customer)
	customerService.SaveCustomer(Customer)
	json.NewEncoder(rw).Encode("Created Succesfully !")
	json.NewEncoder(rw).Encode(Customer)
}

func UpdateCustomer(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Customer models.Customer
	json.NewDecoder(r.Body).Decode(&Customer)
	customerService.UpdateCustomer(Customer)
	json.NewEncoder(rw).Encode(" updated by PUT method !")
	json.NewEncoder(rw).Encode(Customer)
}



