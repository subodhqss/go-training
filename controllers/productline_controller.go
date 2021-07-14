package controllers

import (
	"encoding/json"
	"net/http"


	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var productlineService = service.ProductlineService(repository.NewProdLineRepo())

func GetProductline(rw http.ResponseWriter, r *http.Request) {
	data := productlineService.PrintProductline()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func CreateProductline(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Productline models.Productline
	json.NewDecoder(r.Body).Decode(&Productline)
	productlineService.SaveProductline(Productline)
	json.NewEncoder(rw).Encode("Productline is Created Succesfully !")
	json.NewEncoder(rw).Encode(Productline)
}

func UpdateProductline(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Productline models.Productline
	json.NewDecoder(r.Body).Decode(&Productline)
	productlineService.UpdateProductline(Productline)
	json.NewEncoder(rw).Encode("Productline is updated by PUT method !")
	json.NewEncoder(rw).Encode(Productline)
}

func UpdatePl(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Productline models.Productline
	json.NewDecoder(r.Body).Decode(&Productline)
	productlineService.UpdatePl(Productline)
	json.NewEncoder(rw).Encode("Productline is updated by PATCH method !")
	json.NewEncoder(rw).Encode(Productline)
}
