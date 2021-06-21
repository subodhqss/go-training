package models

type Employee struct {
	Name            string `json:"Employee_Name"`
	Add             string `json:"Address_Of_Employee"`
	Designation     string `json:"Designation_of_Employee"`
	Company_name    string `json:"Company_Name"`
	Date_of_joining string `json:"joining_date"`
	EmpCode         string `json:"Employee_code"`
}
