package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type InvoiceFactory struct{}

func create(attributes map[string]interface{}) *models.Invoice {
	return &models.Invoice{
		Id:         attributes["id"].(int),
		CustomerId: attributes["custumer_id"].(int),
		Total:      attributes["total"].(string),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
