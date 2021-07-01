package models

type Employee struct {
	EmployeeNumber int    `gorm:"column:employeeNumber" json:"employee_number"`
	LastName       string `gorm:"column:lastName" json:"last_name"`
	FirstName      string `gorm:"column:firstName" json:"first_name"`
	Extension      string `gorm:"column:extension"`
	Email          string `gorm:"column:email"`
	OfficeCode     string `gorm:"column:officeCode" json:"office_code"  `
	ReportTo       int    `gorm:"column:reportsTo" json:"report_to"`
	JobTitle       string `gorm:"column:jobTitle" json:"job_title"`
}

// type Offices struct {
// 	officeCode   string
// 	phone        string
// 	addressLine1 string
// 	addressLine2 string
// 	state        string
// 	country      string
// 	postalCode   string
// 	territory    string
// }