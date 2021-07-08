package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	service "github.com/subodhqss/go-training/services"
)

var productService = service.NewProductService(repository.NewProRepo())

func GetProducts(rw http.ResponseWriter, r *http.Request) {
	data := productService.PrintProduct()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func GetProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["cid"]
	data := productService.PrintProductCode(vars)
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
}
func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Product models.Product
	json.NewDecoder(r.Body).Decode(&Product)
	productService.EditProduct(Product)

}
