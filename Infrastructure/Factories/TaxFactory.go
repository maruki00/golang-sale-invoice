package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type TaxFactory struct{}

func (pf *TaxFactory) create(attributes map[string]interface{}) *models.Tax {
	return &models.Tax{
		Id:        attributes["id"].(int),
		Name:      attributes["name"].(string),
		Slug:      attributes["slug"].(string),
		Status:    attributes["status"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
