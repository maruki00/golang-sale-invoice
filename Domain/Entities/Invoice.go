package models

import "time"

type Invoice struct {
	Id         int       `jaon: id`
	CustomerId int       `jaon: custumer_id`
	Total      string    `jaon: total`
	CreatedAt  time.Time `jaon: created_at`
	UpdatedAt  time.Time `jaon: updated_at`
}
