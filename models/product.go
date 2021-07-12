
package models

type Product struct {
	ProductCode        string  `gorm:"column:productCode" json:"product_Code"`
	ProductName        string  `gorm:"column:productName" json:"product_Name"`
	ProductLine        string  `gorm:"column:productLine" json:"product_Line"`
	ProductScale       string  `gorm:"column:productScale" json:"product_Scale"`
	ProductVendor      string  `gorm:"column:productVendor" json:"product_Vendor"`
	ProductDescription string  `gorm:"column:productDescription" json:"product_Description"`
	QuantityInStock    int     `gorm:"column:quantityInStock" json:"quantity_In_Stock"`
	BuyPrice           float64 `gorm:"column:buyPrice" json:"buy_Price"`
	MSRP               float32 `gorm:"column:MSRP" json:"MSRP"`
	ProductlineDetails []Productline `gorm:"foreignKey:ProductLine;references:ProductLine" json:"Payment_details"`

}