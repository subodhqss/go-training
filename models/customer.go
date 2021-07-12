package models

type Customer struct {
	CustomerNumber         int       `gorm:"column:customerNumber" json:"customer_number"`
	CustomerName           string    `gorm:"column:customerName" json:"customer_name"`
	Phone                  string    `gorm:"column:phone" json:"phon_e"`
	SalesRepEmployeeNumber string    `gorm:"column:salesRepEmployeeNumber" json:"sales_refemp_num"`
	ContactFirstname       string    `gorm:"column:contactFirstName" json:"contact_firstname"`
	City                   string    `gorm:"column:city" json:"city"`
	State                  string    `gorm:"column:state" json:"stat_e"`
	Country                string    `gorm:"column:country" json:"countr_y"`
	AddressLine1           string    `gorm:"column:addressLine1" json:"address_line"`
	PostalCode             string    `gorm:"column:postalCode" json:"postal_code"`
	PaymentDetails         []Payment `gorm:"foreignKey:CustomerNumber;references:CustomerNumber"json:"Payment_details"`
}
