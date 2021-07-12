package models

type Order struct {
	OrderNumber    int    `gorm:"column:orderNumber" json:"Order_Number"`
	OrderDate      string `gorm:"column:orderDate" json:"order_date"`
	RequiredDate   string `gorm:"column:requiredDate" json:"required_date"`
	ShippedDate    string `gorm:"column:shippedDate" json:"shipped_date"`
	Status         string `gorm:"column:status " json:"status"`
	CustomerNumber int    `gorm:"column:customerNumber" json:"customer_Number"`
	

}
