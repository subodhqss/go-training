package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var productService = services.ProductService(repository.NewProdRepo())

func GetProductDetails(rw http.ResponseWriter, r *http.Request) {
	data := productService.PrintProduct()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func GetProductCode(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["productCode"]
	data := productService.PrintProductCode(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}


