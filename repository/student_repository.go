package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type StudentRepository interface {
	PrintStudent() []*models.Student
	SaveStudent(models.Student) *models.Student
	UpdateStudent(models.Student) *models.Student
	Deletestudent(RollNo int64) *models.Student
}

func NewStuRepo() StudentRepository {
	return &studentRepo{}
}

type studentRepo struct{}

func (tr *studentRepo) PrintStudent() []*models.Student {

	var student []*models.Student
	result := gormDB.Limit(10).Find(&student)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	return student
}
func (tr *studentRepo) SaveStudent(student models.Student) *models.Student {

	result := gormDB.Create(&student)
	if err := result.Error; err != nil {
		log.Print("Error in Save all records")
	}
	return &student

}
func (tr *studentRepo) UpdateStudent(student models.Student) *models.Student {
	result := gormDB.Model(&models.Student{}).Where("RollNo = ?", student.RollNo).Updates(models.Student{
		Name:         student.Name,
		CouseName:    student.CouseName,
		Email:        student.Email,
		MobileNumber: student.MobileNumber,
	})
	if err := result.Error; err != nil {
		log.Print("Error in update all records")
	}
	fmt.Println(&student)
	return &student

}
func (tr *studentRepo) Deletestudent(RollNo int64) *models.Student {
	result := gormDB.Delete(&models.Student{}, RollNo)
	if err := result.Error; err != nil {
		log.Print("Error in Delete  records")
	}
	return &models.Student{}

}
