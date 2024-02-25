package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"time"
)

type ProductSuplierFactory struct {
}

func (pf *ProductSuplierFactory) create(attributes map[string]interface{}) *models.ProductSupplier {
	return &models.ProductSupplier{
		Id:         attributes["id"].(int),
		ProductId:  attributes["product_id"].(int),
		SupplierId: attributes["supplier_id"].(int),
		Price:      attributes["price"].(float32),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
