package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

type UnitFactory struct{}

func (pf *UnitFactory) create(attributes map[string]interface{}) *entities.Unit {
	return &entities.Unit{
		Name:      attributes["name"].(string),
		Slug:      attributes["slug"].(string),
		Status:    attributes["status"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
