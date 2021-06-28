package controllers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employees", GetEmploye).Methods("GET")
	router.HandleFunc("/saveEmployee", CreateEmployee).Methods("POST")
	router.HandleFunc("/updateEmployee", UpdateEmployee).Methods("PUT")


	return router
}
