package controllers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", GetTest)
	router.HandleFunc("/employee", GetEmployee)
	router.HandleFunc("/saveemployee", CreateEmployee).Methods("POST")
	router.HandleFunc("/student", GetStudent)
	router.HandleFunc("/savestudent", CreateStudent).Methods("POST")
	router.HandleFunc("/updatestudent", UpdateStudent).Methods("PUT")
	router.HandleFunc("/deleteRecord/{RollNo}", Deletestudent).Methods("DELETE")
	return router
}
