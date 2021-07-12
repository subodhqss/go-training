package controllers

import (
	"github.com/gorilla/mux")


func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//employee model methods
	router.HandleFunc("/employees", GetEmploye).Methods("GET")
	router.HandleFunc("/employee/{eid}", GetEmployeId).Methods("GET")
	router.HandleFunc("/save/Employee", CreateEmployee).Methods("POST")
	router.HandleFunc("/update/Employee", UpdateEmployee).Methods("PUT")
	router.HandleFunc("/update/Employee", Update).Methods("PATCH")
	router.HandleFunc("/delete/{eid}", DeleteEmployee).Methods("DELETE")

	//office model methods
	router.HandleFunc("/offices", GetOffices).Methods("GET")
	router.HandleFunc("/office/{code}",GetOffice).Methods("GET")
	router.HandleFunc("/save/Office", CreateOffice).Methods("POST")
	router.HandleFunc("/update/Office", UpdateOffice).Methods("PUT")

	//customer model methods
	router.HandleFunc("/customers", GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{code}",GetCustomer).Methods("GET")
	router.HandleFunc("/save/Customer", CreatCustomer).Methods("POST")
	router.HandleFunc("/update/Customer", UpdateCustomer).Methods("PUT")

	//orderDetails model methods
	router.HandleFunc("/orderDetails", GetOrderDetails).Methods("GET")
	router.HandleFunc("/orderDetail/{code}",GetOrderDetail).Methods("GET")
	router.HandleFunc("/save/OrderDetails", CreatOrderDetail).Methods("POST")
	router.HandleFunc("/update/OrderDetails", UpdateOrderDetail).Methods("PUT")
	
	//products model methods
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/product/{code}",GetProduct).Methods("GET")
	router.HandleFunc("/save/Product", CreateProduct).Methods("POST")
	router.HandleFunc("/update/Product", UpdateProduct).Methods("PUT")
	
	//payment model methods
	router.HandleFunc("/payments", GetPayments).Methods("GET")
	router.HandleFunc("/payment/{code}",GetPayment).Methods("GET")
	router.HandleFunc("/save/Payment", CreatePayment).Methods("POST")
	router.HandleFunc("/update/Payment", UpdatePayment).Methods("PUT")

	//order model methods
	router.HandleFunc("/orders", GetOrders).Methods("GET")
	router.HandleFunc("/order/{code}",GetOrder).Methods("GET")
	router.HandleFunc("/save/Order", CreateOrder).Methods("POST")
	router.HandleFunc("/update/Order", UpdateOrder).Methods("PUT")

	//Productlines model methods
	router.HandleFunc("/productline", GetProductlines).Methods("GET")
	router.HandleFunc("/productline/{code}",GetProductline).Methods("GET")
	router.HandleFunc("/save/Productline", CreateProductline).Methods("POST")
	router.HandleFunc("/update/Productline", UpdateProductline).Methods("PUT")

	return router
}
