package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type CustomerFactory struct{}

func create(attributes map[string]interface{}) *models.Customer {
	return &models.Customer{
		Id:             attributes["Id"].(int),
		Name:           attributes["Name"].(string),
		Mobile:         attributes["Mobile"].(string),
		Address:        attributes["Address"].(string),
		Email:          attributes["Email"].(string),
		Details:        attributes["Details"].(string),
		PreviousBlance: attributes["PreviousBlance"].(string),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
