package models

type Payment struct {
	CustomerNumber int    `gorm:"column:customerNumber" json:"customer_Number"`
	CheckNumber    string `gorm:"column:checkNumber " json:"check_Number "`
	PaymentDate    string `gorm:"column:paymentDate" json:"payment_date"`

	Amount float64 `gorm:"column:amount" json:"amount"`
}
