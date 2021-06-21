package models

type Employee struct {
	Name            string `json:"employee_Name"`
	Add             string `json:"address_Of_Employee"`
	Designation     string `json:"designation_of_Employee"`
	Company_name    string `json:"company_Name"`
	Date_of_joining string `json:"joining_date"`
	EmpCode         string `json:"employee_code"`
}
