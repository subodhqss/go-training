package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	service "github.com/subodhqss/go-training/services"
)

var officeService = service.NewOfficeService(repository.NewOfficeRepo())

func GetOfficee(rw http.ResponseWriter, r *http.Request) {
	data := officeService.PrintOffice()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)

}
func GetOffice(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["officeCode"]
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
}
func UpdateOffice(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Office models.Office
	json.NewDecoder(r.Body).Decode(&Office)
	officeService.EditOffice(Office)

}
