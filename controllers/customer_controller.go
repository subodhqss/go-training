package controllers

import (
	"encoding/json"
	"net/http"

	// "github.com/gorilla/mux"
	// "github.com/subodhqss/go-training/models"
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