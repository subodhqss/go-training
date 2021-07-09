package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var ProductlineService = service.NewProductlineService(repository.NewPrlRepo())
 
func GetProductlines(rw http.ResponseWriter, r *http.Request) {
	data := ProductlineService.PrintProductline()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func GetProductline(rw http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)["code"]
	data := ProductlineService.PrintProductlineId(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreateProductline(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Productline models.Productline
	json.NewDecoder(r.Body).Decode(&Productline)
	ProductlineService.SaveProductline(Productline)
	json.NewEncoder(rw).Encode("Created Succesfully !")
	json.NewEncoder(rw).Encode(Productline)
}

func UpdateProductline(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Productline models.Productline
	json.NewDecoder(r.Body).Decode(&Productline)
	ProductlineService.UpdateProductline(Productline)
	json.NewEncoder(rw).Encode(" updated by PUT method !")
	json.NewEncoder(rw).Encode(Productline)
}



