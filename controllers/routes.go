package controllers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Employee", GetEmploye)
	router.HandleFunc("git/{code}", GetOffice)
	router.HandleFunc("/Employee/{eid}", GetEmployeId)
	router.HandleFunc("/saveemployee", CreateEmployee).Methods("POST")
	router.HandleFunc("/editemployee",UpdateEmployee).Methods("PUT")
	router.HandleFunc("/removeemployee/{eid}",DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/patchemployee",PatchEmployee).Methods("PATCH")


	return router
}
