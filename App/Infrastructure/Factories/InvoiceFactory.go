package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

type InvoiceFactory struct{}

func (pf *InvoiceFactory) create(attributes map[string]interface{}) *entities.Invoice {
	return &entities.Invoice{
		Id:         attributes["id"].(int),
		CustomerId: attributes["custumer_id"].(int),
		Total:      attributes["total"].(string),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
