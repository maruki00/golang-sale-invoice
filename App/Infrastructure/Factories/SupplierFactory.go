package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type SupplierFactory struct{}

func (pf *SupplierFactory) create(attributes map[string]interface{}) *models.Supplier {
	return &models.Supplier{
		Id:              attributes["id"].(int),
		Name:            attributes["name"].(string),
		Mobile:          attributes["mobile"].(string),
		Address:         attributes["address"].(string),
		Details:         attributes["details"].(string),
		PreviousBalance: attributes["previous_balance"].(string),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
