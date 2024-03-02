package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type PurchaseFactory struct {
}

func (pf *PurchaseFactory) create(attributes map[string]interface{}) *models.Purchase {
	return &models.Purchase{
		Id:         attributes["id"].(int),
		SupplierId: attributes["supplier_id"].(int),
		Date:       attributes["date"].(time.Time),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
