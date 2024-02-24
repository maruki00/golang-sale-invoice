package models

import "time"

type Unit struct {
	Name      string    `json:name`
	Slug      string    `json:slug`
	Status    string    `json:status`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}
