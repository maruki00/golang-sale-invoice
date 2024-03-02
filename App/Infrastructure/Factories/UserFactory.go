package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

type UserFactory struct {
}

func (pf *UserFactory) create(attributes map[string]interface{}) *entities.User {
	return &entities.User{
		Id:              attributes["id"].(int),
		FirstName:       attributes["f_name"].(string),
		LastName:        attributes["l_name"].(string),
		Email:           attributes["email"].(string),
		Image:           attributes["image"].(string),
		EmailVerifiedAt: attributes["email_verified_at"].(string),
		Password:        attributes["password"].(string),
		RememberToken:   attributes["remember_token"].(string),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
