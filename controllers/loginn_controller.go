// package controllers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/subodhqss/go-training/models"
// 	"github.com/subodhqss/go-training/repository"
// 	service "github.com/subodhqss/go-training/services"
// )


// var loginService = service.NewLoginService(repository.NewLogRepo())


// func LoginEmployee(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Set("Content-Type", "application/json")
// 	var Login models.Login
// 	json.NewDecoder(r.Body).Decode(&Login)
// 	loginService.LoginEmployee(Login)
// 	jsonData, _ := json.Marshal(Login)
// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.WriteHeader(200)
// 	rw.Write(jsonData)
	

// 	//json.NewEncoder(rw).Encode("Created....")
// 	//json.NewEncoder(rw).Encode(Login)
// }

// // func LoginEmployee(rw http.ResponseWriter, r *http.Request) {
// // 	data := loginService.LoginEmployee(models.Login{})
// // 	jsonData, _ := json.Marshal(data)
// // 	rw.Header().Set("Content-Type", "application/json")
// // 	rw.WriteHeader(200)
// // 	rw.Write(jsonData)

// // }


	
// func Register(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Set("Content-Type", "application/json")
// 	var Login models.Employee
// 	json.NewDecoder(r.Body).Decode(&Login)
// 	employeService.SaveEmployee(Login)
// 	json.NewEncoder(rw).Encode("Created....")
// 	json.NewEncoder(rw).Encode(Login)
// }

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	service "github.com/subodhqss/go-training/services"
)

var loginService = service.NewLoginService(repository.NewLogRepo())

func Login(rw http.ResponseWriter, r *http.Request) {
	// data := loginService.LoginEmployee("email")
	data := "test"
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}
func LoginEmployee(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var login models.Login
	json.NewDecoder(r.Body).Decode(&login)
	data := loginService.LoginEmployee(&login)
	json.NewEncoder(rw).Encode(data)
}