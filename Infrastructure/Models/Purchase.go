package models

import "time"

type Purchase struct {
	Id         int       `json:id`
	SupplierId int       `json:supplier_id`
	Date       time.Time `json:date`
	CreatedAt  time.Time `json:created_at`
	UpdatedAt  time.Time `json:updated_at`
}
