package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	service "github.com/subodhqss/go-training/services"
)

type Credentials struct{
	Email string `json:"email"`
	Password string `json:"password"`

}


var loginService = service.NewLoginService(repository.NewLogRepo())

func Login(rw http.ResponseWriter, r *http.Request) {
	data := loginService.PrintLogin()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func Home(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Login models.Employee
	json.NewDecoder(r.Body).Decode(&Login)
	employeService.SaveEmployee(Login)
	json.NewEncoder(rw).Encode("Created....")
	json.NewEncoder(rw).Encode(Login)
}
