package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

type CustomerFactory struct{}

func (pf *CustomerFactory) create(attributes map[string]interface{}) *entities.Customer {
	return &entities.Customer{
		Id:             attributes["id"].(int),
		Name:           attributes["name"].(string),
		Mobile:         attributes["mobile"].(string),
		Address:        attributes["address"].(string),
		Email:          attributes["email"].(string),
		Details:        attributes["details"].(string),
		PreviousBlance: attributes["previous_balance"].(string),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
