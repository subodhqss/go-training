package models

type Customer struct {
	CustomerNumber int       `gorm:"column:customerNumber" json:"customer_number"`
	CustomerName   string    `gorm:"column:customerName" json:"customer_name"`
	Phone          string    `gorm:"column:phone"`
	AddrressLine1  string    `gorm:"column:addressLine1" json:"address_line"`
	City           string    `gorm:"column:city"`
	State          string    `gorm:"column:state"`
	PostalCode     string    `gorm:"column:postalCode" json:"postal_code"`
	Country        string    `gorm:"column:country"`
	PaymentDetails []Payment `gorm:"foreignKey:CustomerNumber;references:CustomerNumber" json:"Payment_details"`
}
