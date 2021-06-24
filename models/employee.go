package models

type Employee struct {
	// gorm.Model
	EmployeeNumber int    `gorm:"column:employeeNumber"json:"employee_number"`
	LastName       string `gorm:"column:lastName"json:"last_name"`
	FirstName      string `gorm:"column:firstName"json:"first_name"`
	Extension      string `gorm:"column:extension"`
	Email          string `gorm:"column:email"`
	OfficeCode     string `gorm:"column:officeCode"json:"office_code"`
	ReportTo       int    `gorm:"column:reportTo"json:"report_to"`
	JobTitle       string `gorm:"column:jobTitle"json:"job_title"`

	// Name            string `json:"employee_name"`
	// Add             string `json:"address_of_employee"`
	// Designation     string `json:"designation_of_Employee"`
	// Company_name    string `json:"company_name"`
	// Date_of_Joining string `json:"joining_date"`
	// EmpCode         string `json:"employee_code"`

}
