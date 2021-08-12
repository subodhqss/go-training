package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Hello)
	router.HandleFunc("/test", GetTest)
	
	employeeSubRoute := mux.NewRouter().PathPrefix("/employees").Subrouter().StrictSlash(true)
	employeeSubRoute.Methods(http.MethodGet).Path("/{employeeNumber}").HandlerFunc(GetEmployeId)
	employeeSubRoute.Methods(http.MethodDelete).Path("/delete/{employeeNumber}").HandlerFunc(DeleteEmployee)
	employeeSubRoute.Methods(http.MethodGet).HandlerFunc(GetEmploye)
	router.PathPrefix("/employees").Handler(employeeSubRoute)
	// employeeSubRoute.Methods(http.MethodGet).HandlerFunc(GetEmployeId)
	// employeeSubRoute.Methods(http.MethodPost).HandlerFunc(AddEmployee)
	//  router.HandleFunc("/saveEmployee", CreateEmployee).Methods("POST")
	// router.HandleFunc("/employees", GetEmploye).Methods("GET")
	// router.HandleFunc("/employee/{eid}", GetEmployeId).Methods("GET")
	// router.HandleFunc("/updateEmployee", UpdateEmployee).Methods("PUT")
	// router.HandleFunc("/delete/{eid}", DeleteEmployee).Methods(http.MethodDelete)

	//customer model functions
	router.HandleFunc("/customer", GetCustomer).Methods("GET")
	// router.HandleFunc("/customer/{number}", GetCustomerNumber).Methods("GET")
	router.HandleFunc("/savecustomer", CreateCustomer).Methods("POST")
	router.HandleFunc("/editcustomer", UpdateCustomer).Methods("PUT")
	router.HandleFunc("/updatec", UpdateC).Methods("PATCH")

	//offices model functions
	officeSubRoute := mux.NewRouter().PathPrefix("/office").Subrouter().StrictSlash(true)
	officeSubRoute.Methods(http.MethodGet).HandlerFunc(GetOfficee)
	officeSubRoute.Methods(http.MethodGet).Path("/{officeCode}").HandlerFunc(GetOffice)
	router.PathPrefix("/office").Handler(officeSubRoute)

 

	//orderdetail model functions
	orderdetailSubRoute := mux.NewRouter().PathPrefix("/orderdetail").Subrouter().StrictSlash(true)
	orderdetailSubRoute.Methods(http.MethodGet).HandlerFunc(GetDetail)
	orderdetailSubRoute.Methods(http.MethodPost).Path("/{savedetail}").HandlerFunc(CreateDetail)
	router.PathPrefix("/orderdetail").Handler(orderdetailSubRoute)

	// //product model functions
	// productdetailSubRoute := mux.NewRouter().PathPrefix("/product").Subrouter().StrictSlash(true)
	// productdetailSubRoute.Methods(http.MethodGet).HandlerFunc(GetProductDetails)
	// productdetailSubRoute.Methods(http.MethodGet).Path("/{productCode}").HandlerFunc(GetProductCode)
	// router.PathPrefix("/product").Handler(productdetailSubRoute)

//productline model functions
// router.HandleFunc("/productline", GetProductline).Methods("GET")
// router.HandleFunc("/saveproductline", CreateProductline).Methods("POST")

	// productline model functions
	productlineSubRoute := mux.NewRouter().PathPrefix("/productline").Subrouter().StrictSlash(true)
	productlineSubRoute.Methods(http.MethodGet).HandlerFunc(GetProductline)
	productlineSubRoute.Methods(http.MethodPost).Path("/saveproductline").HandlerFunc(CreateProductline)
	productlineSubRoute.Methods(http.MethodPut).Path("/editproductline").HandlerFunc(UpdateProductline)

	router.PathPrefix("/productline").Handler(productlineSubRoute)

	return router
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>First GoLang Application</h1>")
}
