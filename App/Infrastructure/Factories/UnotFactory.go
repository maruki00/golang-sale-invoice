package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type UnitFactory struct{}

func (pf *UnitFactory) create(attributes map[string]interface{}) *models.Unit {
	return &models.Unit{
		Name:      attributes["name"].(string),
		Slug:      attributes["slug"].(string),
		Status:    attributes["status"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
