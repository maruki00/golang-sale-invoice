package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type ProductFactory struct {
}

func (pf *ProductFactory) create(attributes map[string]interface{}) *models.Product {
	return &models.Product{
		Id:           attributes["Id"].(int),
		Name:         attributes["Name"].(string),
		SerialNumber: attributes["SerialNumber"].(int),
		Model:        attributes["Model"].(string),
		CategoryId:   attributes["CategoryId"].(int),
		SalesPrice:   attributes["SalesPrice"].(float32),
		UnitId:       attributes["UnitId"].(int),
		Image:        attributes["Image"].(string),
		TaxId:        attributes["TaxId"].(int),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
