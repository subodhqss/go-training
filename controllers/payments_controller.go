package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var PaymentService = service.NewPaymentService(repository.NewPayRepo())
 
func GetPayments(rw http.ResponseWriter, r *http.Request) {
	data := PaymentService.PrintPayment()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func GetPayment(rw http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)["code"]
	data := PaymentService.PrintPaymentId(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreatePayment(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Payment models.Payment
	json.NewDecoder(r.Body).Decode(&Payment)
	PaymentService.SavePayment(Payment)
	json.NewEncoder(rw).Encode("Created Succesfully !")
	json.NewEncoder(rw).Encode(Payment)
}

func UpdatePayment(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Payment models.Payment
	json.NewDecoder(r.Body).Decode(&Payment)
	PaymentService.UpdatePayment(Payment)
	json.NewEncoder(rw).Encode(" updated by PUT method !")
	json.NewEncoder(rw).Encode(Payment)
}



