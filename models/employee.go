package models

type Employee struct {
	Name            string `json:"Employee Name"`
	Add             string `json:"Address Of Employee"`
	Designation     string `json:"Designation of Employee"`
	Company_name    string `json:"Company Name"`
	Date_of_joining string `json:"joining date"`
	EmpCode         string `json:"Employee code"`
}
