
package models

type Office struct {
	OfficeCode string `gorm:"column:officeCode" json:"office_code"`
	City       string `gorm:"column:city" json:"cit_y"`
	Phone      string `gorm:"column:phone" json:"phon_e"`
	State      string `gorm:"column:state" json:"stat_e"`
	Country    string `gorm:"column:country" json:"countr_y"`
}
