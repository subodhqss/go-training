package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var orderdetailService = services.OrderdetailService(repository.NewOrderRepo())

func GetDetail(rw http.ResponseWriter, r *http.Request) {
	data := orderdetailService.PrintOrderdetail()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

/*
func GetCustomerNumber(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["cno"]
	data := customerService.PrintCustomerNumber(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
*/
func CreateDetail(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Orderdetail models.Orderdetail
	json.NewDecoder(r.Body).Decode(&Orderdetail)
	orderdetailService.SaveDetail(Orderdetail)
	json.NewEncoder(rw).Encode("Orderdetails is Created Succesfully !")
	json.NewEncoder(rw).Encode(Orderdetail)
}

func UpdateDetail(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Orderdetail models.Orderdetail
	json.NewDecoder(r.Body).Decode(&Orderdetail)
	orderdetailService.UpdateDetail(Orderdetail)
	json.NewEncoder(rw).Encode("Orderdetail is updated by PUT method !")
	json.NewEncoder(rw).Encode(Orderdetail)
}

func UpdateD(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Orderdetail models.Orderdetail
	json.NewDecoder(r.Body).Decode(&Orderdetail)
	orderdetailService.UpdateD(Orderdetail)
	json.NewEncoder(rw).Encode("Orderdetail is updated by PATCH method !")
	json.NewEncoder(rw).Encode(Orderdetail)
}