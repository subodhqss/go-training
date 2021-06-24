package controllers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", GetTest)
	router.HandleFunc("/employee", GetEmployee)
	router.HandleFunc("/saveemployee", CreateEmployee).Methods("POST")
	return router
}
