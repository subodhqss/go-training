package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var officeService = service.NewOfficeService(repository.NewOffRepo())
//Employee model functions 
func GetOffices(rw http.ResponseWriter, r *http.Request) {
	data := officeService.PrintOffice()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func GetOffice(rw http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)["code"]
	data := officeService.PrintOfficeId(vars)
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreateOffice(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Office models.Office
	json.NewDecoder(r.Body).Decode(&Office)
	officeService.SaveOffice(Office)
	json.NewEncoder(rw).Encode("Created Succesfully !")
	json.NewEncoder(rw).Encode(Office)
}

func UpdateOffice(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Office models.Office
	json.NewDecoder(r.Body).Decode(&Office)
	officeService.UpdateOffice(Office)
	json.NewEncoder(rw).Encode(" updated by PUT method !")
	json.NewEncoder(rw).Encode(Office)
}



