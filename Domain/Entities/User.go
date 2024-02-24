package models

import "time"

type User struct {
	Id              int          `json:id`
	FirstName       string       `json:f_name`
	LastName        string       `json:l_name`
	Email           string       `json:email`
	Image           string       `json:image`
	EmailVerifiedAt string       `json:email_verified_at`
	Password        string       `json:password`
	RememberToken   string       `json:remember_token`
	CreatedAt       stringe.Time `json:created_at`
	Updated_At      time.Time    `json:updated_at`
}
