package models

type Property struct {
	Id   uint   `json:"id"`
	Name string `json:"name" gorm:"unique"`
}

type ProductProperty struct {
	Id         uint     `json:"id"`
	PropertyID uint     `json:"property_id" gorm:"index:product_property_unq,unique"`
	Property   Property `json:"property"`
	ProductID  uint     `json:"product_id" gorm:"index:product_property_unq,unique"`
	Product    Product  `json:"-"`
	Value      string   `json:"value"`
}
