package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

func CategoryFactory(attributes map[string]interface{}) *entities.Category {
	return &entities.Category{
		Id:        attributes["id"].(int),
		Name:      attributes["name"].(string),
		Slug:      attributes["slug"].(string),
		Status:    attributes["status"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
