package controllers

import (
	"encoding/json"
	"net/http"

	 "github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	 "github.com/subodhqss/go-training/services"
)

var paymentService = services.PaymentService(repository.NewPayRepo())

func GetPayment(rw http.ResponseWriter, r *http.Request) {
	data := paymentService.PrintPayment()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreatePayment(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Payment models.Payment
	json.NewDecoder(r.Body).Decode(&Payment)
	paymentService.SavePayment(Payment)
	json.NewEncoder(rw).Encode("Payment is Created Succesfully !")
	json.NewEncoder(rw).Encode(Payment)
}
