package models

type Employee struct {
	EmployeeNumber int    `gorm:"column:employeeNumber" json:"employee_number"`
	FirstName      string `gorm:"column:firstName" json:"first_name"`
	LastName       string `gorm:"column:lastName" json:"last_name"`
	Extension      string `gorm:"column:extension" json:"extension"`
	Email          string `json:"email"`
	OfficeCode     string `gorm:"column:officeCode" json:"office_Code"`
	ReportTo       int    `gorm:"column:reportTo" json:"reportTo"`
	JobTitle       string `gorm:"column:jobTitle" json:"job_Title"`
}
