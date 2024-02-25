package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type CategoryFactory struct {
}

func (pf *CategoryFactory) create(attributes map[string]interface{}) *models.Category {
	return &models.Category{
		Id:        attributes["Id"].(int),
		Name:      attributes["Name"].(string),
		Slug:      attributes["Slug"].(string),
		Status:    attributes["Status"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
