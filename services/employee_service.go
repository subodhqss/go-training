package service

import (
	"mime/multipart"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/util"
)

type employeService interface {
	PrintEmploye() []*models.Employee
	PrintEmployeId(string) *models.Employee
	//PrintEmployeEmail(string) *models.Employee
	SaveEmployee(models.Employee) *models.Employee
	UpdateEmployee(models.Employee) *models.Employee
	Update(models.Employee) *models.Employee
	//Employeecsv(models.Employee) *models.Employee
	DeleteEmployee(models.Employee, string) *models.Employee
	UpdateImage(string, multipart.File, *multipart.FileHeader, string)
	SaveEmployeeAll([]models.Employee)
}

type empServ struct {
	empRepo repository.EmployeReposiotry
}

func NewEmployeService(empRepo repository.EmployeReposiotry) employeService {
	return &empServ{empRepo: empRepo}
}

//Employee model functions
func (es *empServ) PrintEmploye() []*models.Employee {
	emp := es.empRepo.PrintEmploye()
	return emp
}

func (es *empServ) PrintEmployeId(eid string) *models.Employee {
	emp := es.empRepo.PrintEmployeId(eid)
	return emp
}

/*func (es *empServ) PrintEmployeEmail(email string) *models.Employee {
	emp := es.empRepo.PrintEmployeEmail(email)
	return emp
}*/

func (emp *empServ) SaveEmployee(Employee models.Employee) *models.Employee {

	empId := emp.empRepo.SaveEmployee(Employee)
	return empId
}
func (em *empServ) UpdateEmployee(Employee models.Employee) *models.Employee {
	empId := em.empRepo.UpdateEmployee(Employee)
	return empId
}

func (em *empServ) Update(Employee models.Employee) *models.Employee {
	empId := em.empRepo.Update(Employee)
	return empId
}

func (em *empServ) DeleteEmployee(Employee models.Employee, eid string) *models.Employee {

	empId := em.empRepo.DeleteEmployee(Employee, eid)
	return empId
}
func (em *empServ) UpdateImage(eid string, file multipart.File, header *multipart.FileHeader, host string) {
	imagePath := util.UploadFile(file, header, eid)

	em.empRepo.UpdateImage(host, imagePath, eid)

}

func (emp *empServ) SaveEmployeeAll(Employee []models.Employee) {

	emp.empRepo.SaveEmployeeAll(Employee)

}
