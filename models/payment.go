package models

type Payment struct {
	//OrderNumber int       ` gorm:"column:orderNumber" json:"order_number" `
	CustomerNumber int     `gorm:"column:customerNumber" json:"customer_Number"`
	CheckNumber    string  `gorm:"column:checkNumber " json:"check_Number "`
	Amount         float64 `gorm:"column:amount" json:"amount"`
}
