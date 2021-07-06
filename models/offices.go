package models

type Office struct {
	OfficeCode   string     `gorm:"column:officeCode" json:"office_code"`
	City         string     `gorm:"column:city" json:"city"`
	Phone        string     `gorm:"column:phone" json:"phone"`
	AddressLine1 string     `gorm:"column:addressLine1" json:"addressLine1"`
	State        string     `gorm:"column:state" json:"state"`
	Country      string     `gorm:"column:country" json:"contry"`
	PostalCode   string     `gorm:"column:postalCode" json:"postalCode"`
	Territory   string     `gorm:"column:territory" json:"territory"`
	Employees    []Employee `gorm:"foreignKey:OfficeCode;references:OfficeCode" json:"employees"`
}
