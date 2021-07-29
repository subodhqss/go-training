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
	data := loginService.PrintLogin()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)

	// 	var dbuser Employee

	// 	var credentials Credentials
	// 	err := json.NewDecoder(r.Body).Decode(&credentials)
	// 	if err != nil {
	// 		log.Print("Err1 : ", err)
	// 		rw.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	expectedPassword, ok := user[credentials.Email]

	// 	if !ok || expectedPassword != credentials.Password {
	// 		rw.WriteHeader(http.StatusUnauthorized)
	// 		log.Print("Err2 : ", err)
	// 		return
	// 	}

	// 	expirationTime := time.Now().Add(time.Minute * 5)
	// 	employee := &Employee{
	// 		Email: credentials.Email,
	// 		StandardClaims: jwt.StandardClaims{
	// 			ExpiresAt: expirationTime.Unix(),
	// 		},
	// 	}

	// }
}
func Register(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Login models.Employee
	json.NewDecoder(r.Body).Decode(&Login)
	employeService.SaveEmployee(Login)
	json.NewEncoder(rw).Encode("Created....")
	json.NewEncoder(rw).Encode(Login)
}
