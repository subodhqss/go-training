package models

type Employee struct {
	Name            string `json:"employee_name"`
	Add             string `json:"address_of_employee"`
	Designation     string `json:"Designation of Employee"`
	Company_name    string `json:"Company Name"`
	Date_of_joining string `json:"joining date"`
	EmpCode         string `json:"Employee code"`
}
