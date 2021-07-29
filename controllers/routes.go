package controllers

import "github.com/gorilla/mux"

//log.Fatal(http.ListenAndServe())

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/Employee", GetEmploye).Methods("GET")
	router.HandleFunc("/employee/{eid}", GetEmployeId).Methods("GET")

	//router.HandleFunc("/employee/{eid}", GetEmployeEmail).Methods("GET")
	
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
	router.HandleFunc("/saveOrderDetails", CreateOrderDetail).Methods("POST")
	router.HandleFunc("/updateOrderDetails", UpdateOrderDetail).Methods("PUT")

	//products

	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/product/{code}", GetProduct).Methods("GET")
	router.HandleFunc("/saveproduct", CreateProduct).Methods("POST")
	router.HandleFunc("/updateproduct", UpdateProduct).Methods("PUT")

	//payments
	router.HandleFunc("/payments", GetPayments).Methods("GET")
	router.HandleFunc("/payment/{code}", GetPayment).Methods("GET")
	router.HandleFunc("/savepayment", CreatePayment).Methods("POST")
	router.HandleFunc("/updatepayment", UpdatePayment).Methods("PUT")

	//orders
	router.HandleFunc("/orders", GetOrders).Methods("GET")
	router.HandleFunc("/order/{code}", GetOrder).Methods("GET")
	router.HandleFunc("/saveorder", CreateOrder).Methods("POST")
	router.HandleFunc("/updateorder", UpdateOrder).Methods("PUT")

	//productlines
	router.HandleFunc("/productline", GetProductlines).Methods("GET")
	router.HandleFunc("/productline/{code}", GetProductline).Methods("GET")
	router.HandleFunc("/saveproductline", CreateProductline).Methods("POST")
	router.HandleFunc("/updateproductline", UpdateProductline).Methods("PUT")

	//login-api

	router.HandleFunc("/login", Login).Methods("GET")
	router.HandleFunc("/login", Register).Methods("POST")
	// router.HandleFunc("/welcome", Welcome).Methods("POST")


	// router.HandleFunc("/login", Sigin).Methods("GET")
	// router.HandleFunc("/login", Signup).Methods("POST")
	

	return router

}
