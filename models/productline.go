package models

type Productline struct {
	ProductLine     string `gorm:"column:productLine" json:"product_line"`
	TextDescription string `gorm:"column:textDescription" json:"text_description"`
}
