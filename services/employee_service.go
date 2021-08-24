package service

import (
	"encoding/csv"
	"fmt"

	"os"

	"mime/multipart"

	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/util"
)

type employeService interface {
	PrintEmploye() []*models.Employee
	PrintEmployeId(string) *models.Employee
	SaveEmployee(models.Employee) *models.Employee
	SaveEmployeeCSV(models.Employee) *models.Employee
	UpdateEmployee(models.Employee) *models.Employee
	Update(models.Employee) *models.Employee
	DeleteEmployee(models.Employee, string) *models.Employee
	UpdateImage(string, multipart.File, *multipart.FileHeader, string)

	// PrintOfficeId(string) *models.Office
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

func (emp *empServ) SaveEmployee(Employee models.Employee) *models.Employee {

	empId := emp.empRepo.SaveEmployee(Employee)
	return empId
}

func (emp *empServ) SaveEmployeeCSV(Employee models.Employee) *models.Employee {

	csvFile, _ := os.Open("employee.csv")
	fmt.Println(csvFile)
	reader := csv.NewReader(csvFile)
	line, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(line)
	// line0, _ := strconv.ParseInt(line[0], 0, 10)
	// line6, _ := strconv.ParseInt(line[6], 0, 10)

	// Employee = models.Employee{
	// 	EmployeeNumber: int(line0),
	// 	LastName:       line[1],
	// 	FirstName:      line[2],
	// 	Extension:      line[3],
	// 	Email:          line[4],
	// 	OfficeCode:     line[5],
	// 	ReportsToId:    int(line6),
	// 	JobTitle:       line[7],
	// 	ProfileImage:   line[8],
	// }

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
