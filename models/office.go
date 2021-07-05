package models

type Office struct {
	OfficeCode   string `gorm:"column:officeCode" json:"office_code"`
	City         string `gorm:"column:city"`
	Phone        string `gorm:"column:phone"`
	AddressLine1 string `gorm:"column:addressLine1" json:"address_line1"`
	AddressLine2 string `gorm:"column:addressLine2" json:"address_line2"`
	State        string `gorm:"column:state"`
	Country      string `gorm:"column:country"`
	PostalCode   string `gorm:"column:postalCode" json:"postal_code"`
	Territory    string `gorm:"column:territory"`
}
