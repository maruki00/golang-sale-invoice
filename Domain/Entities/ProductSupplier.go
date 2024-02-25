package models

import "time"

type ProductSupplier struct {
	Id         int       `json: id`
	ProductId  int       `json: product_id`
	SupplierId int       `json: supplier_id`
	Price      float32   `json: price`
	CreatedAt  time.Time `json: created_at`
	UpdatedAt  time.Time `json: updated_at`
}
