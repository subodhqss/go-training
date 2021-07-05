package models

type Office struct {
	OfficeCode string `gorm:"column:officeCode" json:"office_code"`
	City       string `gorm:"column:city" json:"city_"`
	Phone      string `gorm:"column:phone" json:"phone_"`
	State      string `gorm:"column:state" json:"state_"`
	Country    string `gorm:"column:country" json:"contry_"`
}
