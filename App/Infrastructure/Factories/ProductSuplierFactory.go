package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

type ProductSuplierFactory struct {
}

func (pf *ProductSuplierFactory) create(attributes map[string]interface{}) *entities.ProductSupplier {
	return &entities.ProductSupplier{
		Id:         attributes["id"].(int),
		ProductId:  attributes["product_id"].(int),
		SupplierId: attributes["supplier_id"].(int),
		Price:      attributes["price"].(float32),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
