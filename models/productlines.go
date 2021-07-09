package models

type Productline struct {
	ProductLine     string `gorm:"column:productLine" json:"Product_Line"`
	TextDescription string `gorm:"column:textDescription" json:"Text_Description"`
	HtmlDescription string `gorm:"column:htmlDescription" json:"Html_Description"`
	Image           string `gorm:"column:image" json:"image"`
}
