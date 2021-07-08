package controllers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Employee", GetEmploye).Methods("GET")
	router.HandleFunc("/employee/{eid}", GetEmployeId).Methods("GET")
	router.HandleFunc("/saveemployee", CreateEmployee).Methods("POST")
	router.HandleFunc("/editemployee", UpdateEmployee).Methods("PUT")
	router.HandleFunc("/delete/{eid}", DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/mergePatch", UpdatePatch).Methods("PATCH")

	//office

	router.HandleFunc("/Office", GetOfficee).Methods("GET")
	router.HandleFunc("/office/{code}", GetOffice).Methods("GET")
	router.HandleFunc("/saveofficee", CreateOffice).Methods("POST")
	router.HandleFunc("/editoffice", UpdateOffice).Methods("PUT")

	//customers
	router.HandleFunc("/Customer", GetCustomer).Methods("GET")
	router.HandleFunc("/customer/{cid}", GetCustomers).Methods("GET")
	router.HandleFunc("/savecustomer", CreateCustomer).Methods("POST")
	router.HandleFunc("/editcustomer", UpdateCustomer).Methods("PUT")

	//orderDetails
	router.HandleFunc("/orderdetails", GetOrderDetails).Methods("GET")
	router.HandleFunc("/orderDetail/{code}", GetOrderDetail).Methods("GET")
	router.HandleFunc("/saveOrderDetails", CreatOrderDetail).Methods("POST")
	router.HandleFunc("/updateOrderDetails", UpdateOrderDetail).Methods("PUT")

	//products
	//router.HandleFunc("/productdetails", GetProductDetails).Methods("GET")
	//router.HandleFunc("/productDetail/{code}", GetProductDetail).Methods("GET")
	//router.HandleFunc("/saveproductDetails", CreatproductDetail).Methods("POST")
	//router.HandleFunc("/updateproductDetails", UpdateproductDetail).Methods("PUT")


	return router
}
