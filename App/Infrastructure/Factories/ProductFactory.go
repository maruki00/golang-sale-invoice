package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

type ProductFactory struct {
}

func (pf *ProductFactory) create(attributes map[string]interface{}) *entities.Product {
	return &entities.Product{
		Id:           attributes["id"].(int),
		Name:         attributes["name"].(string),
		SerialNumber: attributes["serial_number"].(int),
		Model:        attributes["model"].(string),
		CategoryId:   attributes["category_id"].(int),
		SalesPrice:   attributes["sales_price"].(float32),
		UnitId:       attributes["unit_id"].(int),
		Image:        attributes["image"].(string),
		TaxId:        attributes["tax_id"].(int),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
