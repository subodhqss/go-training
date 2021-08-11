package repository
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/go-training/models"
)
func TestSaveEmployee1(t *testing.T){
	repo := NewEmpRepo();
	employeeDummy := &models.Employee{
		EmployeeNumber: 109,
		LastName:  "kaur"
		FirstName:  "Neet"
		Extension:   "abcc"
		Email:       "neet@gmail.com"
		ReportsToId:  1002,
		JobTitle:     "Trainee"
	}
	res := repo.SaveEmployee(*employeeDummy)
	employeeDummy.EmployeeNumber = res.EmployeeNumber
	assert.Equal(t,employeeDummy,res)

	employeeDummy1:=models.Employee{
		EmployeeNumber: 999,
	}

	errCheck := repo.SaveEmployee(employeeDummy1)
	assert.Nil(t,errCheck)

}

func TestGetEmployee(t *testing.T) {
	repo := NewEmpRepo()

	var dummyEmployee []*models.Employee
	result := gormDB.Preload("OfficeDetail").Preload("ReportsTo").Limit(10).Find(&dummyEmployee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	res := repo.PrintEmploye()
	assert.Equal(t, dummyEmployee, res)

}

func TestGetEmployeeId(t *testing.T) {
	repo := NewEmpRepo()

	var dummyEmployee *models.Employee
	result := gormDB.Preload("OfficeDetail").Preload("ReportsTo").Where("employeeNumber", 99).Find(&dummyEmployee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	res := repo.PrintEmployeId("99")
	assert.Equal(t, dummyEmployee, res)

	errCheck := repo.PrintEmployeId("jsdasj") //its not working
	assert.NotNil(t, errCheck)

}

func TestUpdateEmployee(t *testing.T) {
	repo := NewEmpRepo()
	employeeDummy := &models.Employee{
		EmployeeNumber: 99,
		LastName:       "Gurneet",
		FirstName:      "Murphy",
		Extension:      "x5800",
		Email:          "dmurphy@classicmodelcars.com",
		OfficeCode:     "1",
		ReportsToId:    1002,
		JobTitle:       "President",
	}

	res := repo.UpdateEmployee(*employeeDummy)
	assert.Equal(t, employeeDummy, res)
	errModel := models.Employee{
		EmployeeNumber: 003,
		LastName:      "Gurneet",
		FirstName:     "Murphyy",
		Extension:     "x5800",
		Email:         "dmurphy@classicmodelcars.com",
		OfficeCode:    "1",
		ReportsToId:   1002,
		JobTitle:      "President",
	}

	errCheck := repo.UpdateEmployee(errModel)
	assert.NotNil(t,errCheck)//not working

}

func TestDelete(t *testing.T){
	repo := NewEmpRepo()

	dummyData:=models.Employee{}

	res:=repo.DeleteEmployee(dummyData,"836674")
	// expResult := gormDB.Where("employeeNumber",836674).Delete(&dummyData)
	assert.Equal(t,res,&dummyData)
}

