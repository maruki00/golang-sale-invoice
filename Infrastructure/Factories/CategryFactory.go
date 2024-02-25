package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type CategoryFactory struct {
}

func (pf *CategoryFactory) create(attributes map[string]interface{}) *models.Category {
	return &models.Category{
		Id:        attributes["id"].(int),
		Name:      attributes["name"].(string),
		Slug:      attributes["slug"].(string),
		Status:    attributes["status"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
