package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var productService = service.ProductService(repository.NewProdRepo())

func GetProduct(rw http.ResponseWriter, r *http.Request) {
	data := productService.PrintProduct()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func CreateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Product models.Product
	json.NewDecoder(r.Body).Decode(&Product)
	productService.SaveProduct(Product)
	json.NewEncoder(rw).Encode("Product is Created Succesfully !")
	json.NewEncoder(rw).Encode(Product)
}

func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Product models.Product
	json.NewDecoder(r.Body).Decode(&Product)
	productService.UpdateProduct(Product)
	json.NewEncoder(rw).Encode("Product is updated by PUT method !")
	json.NewEncoder(rw).Encode(Product)
}

func UpdateP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Product models.Product
	json.NewDecoder(r.Body).Decode(&Product)
	productService.UpdateP(Product)
	json.NewEncoder(rw).Encode("Product is updated by PATCH method !")
	json.NewEncoder(rw).Encode(Product)
}
