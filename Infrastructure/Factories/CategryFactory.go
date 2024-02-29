package factories

import (
	models "saleinvoice/Domain/Entities"
	"time"
)

func CategoryFactory(attributes map[string]interface{}) *models.Category {
	return &models.Category{
		Id:        attributes["id"].(int),
		Name:      attributes["name"].(string),
		Slug:      attributes["slug"].(string),
		Status:    attributes["status"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
