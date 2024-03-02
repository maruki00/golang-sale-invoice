package models

import "time"

type Supplier struct {
	Id              int       `json:id`
	Name            string    `json:name`
	Mobile          string    `json:mobile`
	Address         string    `json:address`
	Details         string    `json:details`
	PreviousBalance string    `json:previous_balance`
	CreatedAt       time.Time `json:created_at`
	UpdatedAt       time.Time `json:updated_at`
}
