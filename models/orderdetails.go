package models

type Orderdetail struct {
	OrderNumber     int    `gorm:"column:orderNumber" json:"order_number"`
	ProductCode     string `gorm:"column:productCode" json:"product_code"`
	QuantityOrdered int    `gorm:"column:quantityOrdered" json:"quantity_ordered"`
}
