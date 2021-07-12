package models

type Order struct {
	OrderNumber    int           `gorm:"column:orderNumber" json:"order_Number"`
	OrderDate      string        `gorm:"column:orderDate" json:"order_Date"`
	RequiredDate   string        `gorm:"column:requiredDate" json:"required_Date"`
	ShippedDate    string        `gorm:"column:shippedDate" json:"shipped_Date"`
	Status         string        `gorm:"column:status" json:"status"`
	Comments       string        `gorm:"column:comments" json:"comments"`
	CustomerNumber int           `gorm:"column:customerNumber" json:"customer_Number"`
	OrderDetails   []Orderdetail `gorm:"foreignKey:OrderNumber;references:OrderNumber" json:"Payment_details"`
}
