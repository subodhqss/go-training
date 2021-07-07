package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var OrderDetailService = service.NewOrderDetailService(repository.NewOrddRepo())
 
func GetOrderDetails(rw http.ResponseWriter, r *http.Request) {
	data := OrderDetailService.PrintOrderDetail()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func GetOrderDetail(rw http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)["code"]
	data := OrderDetailService.PrintOrderDetailId(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreatOrderDetail(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var OrderDetail models.Orderdetail
	json.NewDecoder(r.Body).Decode(&OrderDetail)
	OrderDetailService.SaveOrderDetail(OrderDetail)
	json.NewEncoder(rw).Encode("Created Succesfully !")
	json.NewEncoder(rw).Encode(OrderDetail)
}

func UpdateOrderDetail(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var OrderDetail models.Orderdetail
	json.NewDecoder(r.Body).Decode(&OrderDetail)
	OrderDetailService.UpdateOrderDetail(OrderDetail)
	json.NewEncoder(rw).Encode(" updated by PUT method !")
	json.NewEncoder(rw).Encode(OrderDetail)
}



