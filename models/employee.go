package models

type Employee struct {
	EmployeeNumber int       ` gorm:"column:employeeNumber" json:"employee_number" `
	LastName       string    `gorm:"column:lastName" json:"last_name"`
	FirstName      string    `gorm:"column:firstName" json:"first_name"`
	Extension      string    `gorm:"column:extension"`
	Email          string    `gorm:"column:email"`
	// OfficeCode     string    `gorm:"column:officeCode" json:"office_code"  `
	OfficeCode     Offices   `gorm:"foreignKey:OfficeCode; references:OfficeCode"`
	ReportTo       *Employee `gorm:"foreignKey:EmployeeNumber; references:EmployeeNumber" json:"report_to"`
	JobTitle       string    `gorm:"column:jobTitle" json:"job_title"`
}

type Offices struct {
	OfficeCode   string		`gorm:"column:officeCode" json:"office_code"`
	Phone        string		`gorm:"column:phone" json:"phone_"`
	Country      string		`gorm:"column:country" json:"country_"`
}
