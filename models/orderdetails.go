package models

type Orderdetail struct {
	OrderNumber     int    `gorm:"column:orderNumber" json:"OrderNumber"`
	ProductCode     string `gorm:"column:productCode" json:"product_code"`
	ProductDetails  []Product `gorm:"foreignKey:ProductCode;references:ProductCode" json:"Product_details"`
	QuantityOrdered int     `gorm:"column:quantityOrdered" json:"quantityOrdered"`
	PriceEach       float64 `gorm:"column:priceEach" json:"priceEach"`
	OrderLineNumber int     `gorm:"column:orderLineNumber" json:"orderLine_Number"`
}
