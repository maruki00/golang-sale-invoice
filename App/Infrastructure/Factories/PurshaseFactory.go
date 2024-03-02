package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

type PurchaseFactory struct {
}

func (pf *PurchaseFactory) create(attributes map[string]interface{}) *entities.Purchase {
	return &entities.Purchase{
		Id:         attributes["id"].(int),
		SupplierId: attributes["supplier_id"].(int),
		Date:       attributes["date"].(time.Time),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
