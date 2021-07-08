package models

type Product struct {
	//OrderNumber int       ` gorm:"column:orderNumber" json:"order_number" `
	ProductName     string   `gorm:"column:productName" json:"product_Name"`
	ProductCode     string  `gorm:"column:productCode " json:"product_Code "`
	ProductLine 	string     `gorm:"column:productLine" json:"product_line"`
	//ProductScale       float64 `gorm:"column:productScale" json:"product_scale"`
	ProductVendor 	string     `gorm:"column:productVendor" json:"product_vandor"`
	
	
}
