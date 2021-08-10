package models

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmployeeNumber int       ` gorm:"column:employeeNumber" json:"employee_number" `
	LastName       string    `gorm:"column:lastName" json:"last_name"`
	FirstName      string    `gorm:"column:firstName" json:"first_name"`
	Extension      string    `gorm:"column:extension"`
	Email          string    `gorm:"column:email"`
	OfficeCode     string    `gorm:"column:officeCode" json:"office_code"`
	OfficeDetail   Office    `gorm:"foreignKey:OfficeCode;references:OfficeCode" json:"office_detail"`
	ReportsTo      *Employee `gorm:"foreignKey:ReportsToId;references:EmployeeNumber" json:"reports_to"`
	ReportsToId    int       `gorm:"column:reportsTo" json:"reports_to_id"`
	JobTitle       string    `gorm:"column:jobTitle" json:"job_title"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	// uuidStr := uuid.New().String()
	a, _ := rand.Int(rand.Reader, new(big.Int).SetInt64(time.Now().UnixNano()))
	// e.EmployeeNumber = a.String()[13:]
	i, _ := strconv.Atoi(a.String()[13:])
	e.EmployeeNumber = i
	return nil
}
