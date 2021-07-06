package models

type Office struct {
	OfficeCode string     `gorm:"column:officeCode" json:"office_code"`
	City       string     `gorm:"column:city" json:"city"`
	Phone      string     `gorm:"column:phone" json:"phone"`
	State      string     `gorm:"column:state" json:"state"`
	Country    string     `gorm:"column:country" json:"contry"`
	Employees  []Employee `gorm:"foreignKey:OfficeCode;references:OfficeCode" json:"employees"`
}
