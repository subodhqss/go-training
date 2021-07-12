
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var ProductService = service.NewProductService(repository.NewProRepo())
 
func GetProducts(rw http.ResponseWriter, r *http.Request) {
	data := ProductService.PrintProduct()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func GetProduct(rw http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)["code"]
	data := ProductService.PrintProductId(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Product models.Product
	json.NewDecoder(r.Body).Decode(&Product)
	ProductService.SaveProduct(Product)
	json.NewEncoder(rw).Encode("Created >>> !")
	json.NewEncoder(rw).Encode(Product)
}

func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Product models.Product
	json.NewDecoder(r.Body).Decode(&Product)
	ProductService.UpdateProduct(Product)
	json.NewEncoder(rw).Encode(" updated >>> !")
	json.NewEncoder(rw).Encode(Product)
}