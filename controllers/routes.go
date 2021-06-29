package controllers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Employee", GetEmploye)
	router.HandleFunc("/saveemployee", CreateEmployee).Methods("POST")
	router.HandleFunc("/editemployee",UpdateEmployee).Methods("PUT")
	router.HandleFunc("/removeemployee/{eid}",DeleteEmployee).Methods("DELETE")


	return router
}
