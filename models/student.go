package models

type Student struct {
	RollNo       int    `gorm:"column:RollNo"json:"RollNo"`
	Name         string `gorm:"column:Name"json:"Name"`
	CouseName    string `gorm:"column:CouseName"json:"CouseName"`
	Email        string `gorm:"column:email"`
	MobileNumber string `gorm:"column:MobileNumber"json:"MobileNumber"`
	status       string `gorm:"column:status"json:"status"`
}
