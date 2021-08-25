package models

type Product struct {
	ProductCode         string       `gorm:"column:productCode" json:"product_code"`
	ProductName         string       `gorm:"column:productName" json:"product_name"`
	ProductLine         string       `gorm:"column:productLine" json:"product_line"`
	ProductScale        string       `gorm:"column:productScale" json:"product_scale"`
	ProductVendor       string       `gorm:"column:productVendor" json:"product_vendor"`
	ProductDescription  string       `gorm:"column:productDescription" json:"product_description"`
	QuantityInStock     string       `gorm:"column:quantityInStock" json:"quantity_in_stock"`
	BuyPrice            string       `gorm:"column:buyPrice" json:"buy_price"`
	ProductlineDetails []Productline `gorm:"foreignKey:ProductLine;references:ProductLine" json:"Payment_details"`
}