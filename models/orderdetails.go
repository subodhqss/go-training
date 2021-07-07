package models

type OrderDetail struct {
	OrderNumber     int     `gorm:"column:orderNumber" json:"Order_Number"`
	ProductCode     string  `gorm:"column:productCode " json:"produc_tCode "`
	QuantityOrdered int     `gorm:"column:quantityOrdered" json:"quantity_Ordered"`
	PriceEach       float64 `gorm:"column:priceEach" json:"price_Each"`
	OrderLineNumber int     `gorm:"column:orderLineNumber" json:"order_Line_Number"`
}
