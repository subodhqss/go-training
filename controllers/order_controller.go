package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	service "github.com/subodhqss/go-training/services"
)

var OrderService = service.NewOrderService(repository.NewOrdRepo())

func GetOrders(rw http.ResponseWriter, r *http.Request) {
	data := OrderService.PrintOrder()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func GetOrder(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["code"]
	data := OrderService.PrintOrderId(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreateOrder(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Order models.Order
	json.NewDecoder(r.Body).Decode(&Order)
	OrderService.SaveOrder(Order)
	json.NewEncoder(rw).Encode("Created...")
	json.NewEncoder(rw).Encode(Order)
}

func UpdateOrder(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Order models.Order
	json.NewDecoder(r.Body).Decode(&Order)
	OrderService.UpdateOrder(Order)
	json.NewEncoder(rw).Encode(" Updated")
	json.NewEncoder(rw).Encode(Order)
}
