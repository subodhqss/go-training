package models

type Customer struct {
	CustomerNumber         int     `gorm:"column:customerNumber" json:"customer_Number"`
	CustomerName           string  `gorm:"column:customerName" json:"customer_Name"`
	ContactLastName        string  `gorm:"column:contactLastName" json:"contact_LastName"`
	ContactFirstName       string  `gorm:"column:contactFirstName" json:"contact_FirstName"`
	Phone                  string  `gorm:"column:phone" json:"phone"`
	AddressLine1           string  `gorm:"column:addressLine1" json:"address_Line1"`
	City                   string  `gorm:"column:city" json:"city "`
	Country                string  `gorm:"column:country" json:"country"`
	PaymentDetails 		[]Payment `gorm:"foreignKey:CustomerNumber;references:CustomerNumber" json:"Payment_details"`
}
	
