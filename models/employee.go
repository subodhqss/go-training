package models

type Employee struct {
	Name            string `json:"employee_name"`
	Add             string `json:"address_of_employee"`
	Designation     string `json:"Designation_of_Employee"`
	Company_name    string `json:"Company_Name"`
	Date_of_Joining string `json:"joining_date"`
	EmpCode         string `json:"Employee_code"`
}
