package models

type Product struct {
	ProductCode   string `gorm:"column:productCode" json:"product_code"`
	ProductName   string `gorm:"column:productName" json:"product_name"`
	ProductScale  string `gorm:"column:productScale" json:"product_scale"`
	ProductVendor string `gorm:"column:productVendor" json:"product_vendor"`
}
