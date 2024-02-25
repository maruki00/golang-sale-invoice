package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type ProductFactory struct {
}

func (pf *ProductFactory) create(attributes map[string]interface{}) *models.Product {
	return &models.Product{
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
