package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
)

var studentService = services.NewStudentService(repository.NewStuRepo())

func GetStudent(rw http.ResponseWriter, r *http.Request) {
	data := studentService.PrintStu()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)
}

func CreateStudent(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Student models.Student
	json.NewDecoder(r.Body).Decode(&Student)
	studentService.Savestudent(Student)

}

func UpdateStudent(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var Student models.Student
	json.NewDecoder(r.Body).Decode(&Student)
	studentService.Updatestudent(Student)
}

func Deletestudent(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)["RollNo"]
	i1, _ := strconv.ParseInt(vars, 10, 64)
	studentService.Deletestudent(i1)
}
