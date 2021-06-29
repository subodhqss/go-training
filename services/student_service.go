package services

import (
	"fmt"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type StudentService interface {
	PrintStu() []*models.Student
	Savestudent(models.Student) *models.Student
	Updatestudent(models.Student) *models.Student
	Deletestudent(RollNo int64) *models.Student
}

type stuSrv struct {
	// employeeRepo repository.EmployeeRepository
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &stuSrv{studentRepo: studentRepo}
}

func (stu *stuSrv) PrintStu() []*models.Student {
	// employee := stu.employeeRepo.PrintEmployee()
	// test := &models.Test{ID: 1, Message: "Hi SUbodh"}
	student := stu.studentRepo.PrintStudent()
	return student
}

func (stu *stuSrv) Savestudent(Student models.Student) *models.Student {
	// var empId int64 = 0
	// empId := stu.employeeRepo.SaveEmployee(Employee

	return stu.studentRepo.SaveStudent(Student)

}
func (stu *stuSrv) Updatestudent(Student models.Student) *models.Student {
	// var empId int64 = 0
	// empId := stu.employeeRepo.SaveEmployee(Employee

	return stu.studentRepo.UpdateStudent(Student)

}
func (stu *stuSrv) Deletestudent(RollNo int64) *models.Student {
	// var empId int64 = 0
	// empId := stu.employeeRepo.SaveEmployee(Employee
	fmt.Println("In side Service::  ", RollNo)
	return stu.studentRepo.Deletestudent(RollNo)

}
