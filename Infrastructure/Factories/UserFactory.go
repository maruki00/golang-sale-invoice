package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type UserFactory struct {
}

func (pf *UserFactory) create(attributes map[string]interface{}) *models.User {
	return &models.User{
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
