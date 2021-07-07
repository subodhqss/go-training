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
	router.HandleFunc("/offices", GetOffices).Methods("GET")
	router.HandleFunc("/office/{code}",GetOffice).Methods("GET")
	router.HandleFunc("/saveOffice", CreateOffice).Methods("POST")
	router.HandleFunc("/updateOffice", UpdateOffice).Methods("PUT")

	//customer model methods
	router.HandleFunc("/customers", GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{code}",GetCustomer).Methods("GET")
	router.HandleFunc("/saveCustomer", CreatCustomer).Methods("POST")
	router.HandleFunc("/updateCustomer", UpdateCustomer).Methods("PUT")

	return router
}
