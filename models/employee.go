package models

type Employee struct {
	Name            string `json:"name"`
	Add             string `json:"Add"`
	Designation     string `json:"Designation"`
	Company_name    string `json:"Company Name"`
	Date_of_joining string `json:"joining date"`
	EmpCode         string `json:"Employee code"`
}
