package models

type Payment struct {
	CustomerNumber int     `gorm:"column:customerNumber" json:"customer_number"`
	CheckNumber    string  `gorm:"column:checkNumber" json:"check_number"`
	PaymentDate    string  `gorm:"column:paymentDate" json:"payment_date"`
	Amount         float32 `gorm:"column:amount"`
}
