package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type CustomerFactory struct{}

func create(attributes map[string]interface{}) *models.Customer {
	return &models.Customer{
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
