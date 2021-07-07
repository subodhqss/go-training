package models

type Office struct {
	OfficeCode   string `gorm:"column:officeCode" json:"office_code"`
	City         string `gorm:"column:city" json:"cit_y"`
	Phone        string `gorm:"column:phone" json:"phon_e"`
	State        string `gorm:"column:state" json:"stat_e"`
	Country      string `gorm:"column:country" json:"countr_y"`
	AddressLine1 string `gorm:"column:addressLine1" json:"addressLine1"`
	PostalCode   string `gorm:"column:postalCode" json:"postalCode"`
	Territory    string `gorm:"column:territory" json:"territory"`

	Employees []Employee `gorm:"foreignKey:OfficeCode;references:OfficeCode" json:"employees"`
}
