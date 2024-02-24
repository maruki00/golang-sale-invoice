package models

import "time"

type Sale struct {
	Id        int       `json:id`
	InvoiceId int       `json:invoice_id`
	ProductId int       `json:product_id`
	Qty       int       `json:qty`
	Price     float32   `json:price`
	Dis       float32   `json:dis`
	Amount    float32   `json:amount`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}
