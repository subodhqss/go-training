package controllers

import (
	"encoding/json"
	"fmt"
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
	json.NewEncoder(rw).Encode("Created Succesfully !")
	json.NewEncoder(rw).Encode(Product)
}

func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Product models.Product
	json.NewDecoder(r.Body).Decode(&Product)
	ProductService.UpdateProduct(Product)
	json.NewEncoder(rw).Encode(" updated by PUT method !")
	json.NewEncoder(rw).Encode(Product)
}

func ProductImage(rw http.ResponseWriter, r *http.Request) {
	

	rw.Header().Set("Content-Type", "application/json")
	code1 := mux.Vars(r)["code"]

	r.ParseMultipartForm(10 << 20)
	file, header, _ := r.FormFile("myFile")
	defer file.Close()
	fmt.Println("controller")

	ProductService.UpdateImage(code1,file,header,r.Host)
	json.NewEncoder(rw).Encode("product code: "+code1+ " Image updated Succesfully !")

	
}



