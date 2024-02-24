package models

import "time"

type Product struct {
	Id           int       `json:id`
	Name         string    `json:name`
	SerialNumber int       `json:serial_number`
	Model        string    `json:model`
	CategoryId   int       `json:category_id`
	SalesPrice   float32   `json:sales_price`
	UnitId       int       `json:unit_id`
	Image        string    `json:image`
	TaxId        int       `json:tax_id`
	CreatedAt    time.Time `json:created_at`
	UpdatedAt    time.Time `json:updated_at`
}
