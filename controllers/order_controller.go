package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	service "github.com/subodhqss/go-training/services"
)

var orderService = service.OrderService(repository.NewOrdRepo())

func GetOrder(rw http.ResponseWriter, r *http.Request) {
	data := orderService.PrintOrder()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreateOrder(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Order models.Order
	json.NewDecoder(r.Body).Decode(&Order)
	orderService.SaveOrder(Order)
	json.NewEncoder(rw).Encode("Order is Created Succesfully !")
	json.NewEncoder(rw).Encode(Order)
}

func UpdateOrder(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Order models.Order
	json.NewDecoder(r.Body).Decode(&Order)
	orderService.UpdateOrder(Order)
	json.NewEncoder(rw).Encode("Order is updated by PUT method !")
	json.NewEncoder(rw).Encode(Order)
}

func UpdateO(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Order models.Order
	json.NewDecoder(r.Body).Decode(&Order)
	orderService.UpdateO(Order)
	json.NewEncoder(rw).Encode("Order is updated by PATCH method !")
	json.NewEncoder(rw).Encode(Order)
}
