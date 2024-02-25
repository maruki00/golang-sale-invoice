package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type InvoiceFactory struct{}

func create(attributes map[string]interface{}) *models.Invoice {
	return &models.Invoice{
		Id:         attributes["Id"].(int),
		CustomerId: attributes["CustomerId"].(int),
		Total:      attributes["Total"].(string),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
