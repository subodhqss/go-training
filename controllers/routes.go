package controllers

import (
	"github.com/gorilla/mux")


func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//employee model methods
	router.HandleFunc("/employees", GetEmploye).Methods("GET")
	router.HandleFunc("/employee/{eid}", GetEmployeId).Methods("GET")
	router.HandleFunc("/saveEmployee", CreateEmployee).Methods("POST")
	router.HandleFunc("/updateEmployee", UpdateEmployee).Methods("PUT")
	router.HandleFunc("/update", Update).Methods("PATCH")
	router.HandleFunc("/delete/{eid}", DeleteEmployee).Methods("DELETE")

	//office model methods
	router.HandleFunc("/office", GetOfficee).Methods("GET")
	router.HandleFunc("/office/{code}", GetOffice).Methods("GET")
	router.HandleFunc("/saveofficee", CreateOffice).Methods("POST")
	router.HandleFunc("/editoffice", UpdateOffice).Methods("PUT")

	//customer model functions
	router.HandleFunc("/customer", GetCustomer).Methods("GET")
	router.HandleFunc("/customer/{number}", GetCustomer).Methods("GET")
	router.HandleFunc("/savecustomer", CreateCustomer).Methods("POST")
	router.HandleFunc("/editcustomer", UpdateCustomer).Methods("PUT")
	router.HandleFunc("/updatec", Update).Methods("PATCH")





	
	return router
}
