package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

type TaxFactory struct{}

func (pf *TaxFactory) create(attributes map[string]interface{}) *entities.Tax {
	return &entities.Tax{
		Id:        attributes["id"].(int),
		Name:      attributes["name"].(string),
		Slug:      attributes["slug"].(string),
		Status:    attributes["status"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
