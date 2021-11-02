package models

type Product struct {
	Id                uint              `json:"id"`
	Name              string            `json:"name"`
	Count             uint              `json:"count"`
	CategoryID        uint              `json:"category_id"`
	Category          Category          `json:"category"`
	ProductProperties []ProductProperty `json:"product_properties"`
	IsDeleted         bool              `json:"is_deleted"`
}
