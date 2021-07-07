package models

type Orderdetail struct {
	//OrderNumber int       ` gorm:"column:orderNumber" json:"order_number" `
	OrderNumber     int     `gorm:"column:orderNumber" json:"Order_Number"`
	ProductCode     string  `gorm:"column:productCode " json:"product_Code "`
	QuantityOrdered int     `gorm:"column:quantityOrdered" json:"quantity_Ordered"`
	PriceEach       float64 `gorm:"column:priceEach" json:"price_Each"`
	OrderLineNumber int     `gorm:"column:orderLineNumber" json:"orderLine_Number"`
	
	
}
