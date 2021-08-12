package models

type Orderdetail struct {
	OrderNumber     int      `gorm:"column:orderNumber" json:"order_number"`
	ProductCode     string   `gorm:"column:productCode" json:"product_code"`
	ProductDetails  string   `gorm:"foreignKey:ProductCode" json:"Product_details"`
	QuantityOrdered int      `gorm:"column:quantityOrdered" json:"quantity_ordered"`
}