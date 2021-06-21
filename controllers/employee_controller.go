package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)
var employeService = service.NewEmployeService(repository.NewEmpRepo())

func GetEmploye(rw http.ResponseWriter,r *http.Request){
	data:=employeService.PrintEmploye()

	jsonData, _ := json.Marshal(data)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}