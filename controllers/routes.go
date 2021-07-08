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
	router.HandleFunc("/updatec", UpdateC).Methods("PATCH")

	//orderdetail model functions
	router.HandleFunc("/orderdetail", GetDetail).Methods("GET")
	router.HandleFunc("/savedetail", CreateDetail).Methods("POST")
	router.HandleFunc("/editdetail", UpdateDetail).Methods("PUT")
	router.HandleFunc("/updated", UpdateD).Methods("PATCH")
	
	//product model functions
	router.HandleFunc("/product", GetProduct).Methods("GET")
	router.HandleFunc("/saveproduct", CreateProduct).Methods("POST")
	router.HandleFunc("/editproduct", UpdateProduct).Methods("PUT")
	router.HandleFunc("/updatep", UpdateP).Methods("PATCH")

	//payment model functions
	router.HandleFunc("/payment", GetPayment).Methods("GET")
	router.HandleFunc("/savepayment", CreatePayment).Methods("POST")
	router.HandleFunc("/editpayment", UpdatePayment).Methods("PUT")
	router.HandleFunc("/updatepa", UpdatePa).Methods("PATCH")



	











	
	return router
}
