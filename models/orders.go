package models

type Order struct {
	OrderNumber    int    `gorm:"column:orderNumber" json:"order_number"`
	OrderDate      string `gorm:"column:orderDate" json:"order_date"`
	RequiredDate   string `gorm:"column:requiredDate" json:"required_date"`
	ShippedDate    string `gorm:"column:shippedDare" json:"shipped_date"`
	Status         string `gorm:"column:status"`
	CustomerNumber int    `gorm:"column:orderNumber" json:"customer_number"`
}
