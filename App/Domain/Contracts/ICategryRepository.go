package contracts

import (
	entities "saleinvoice/App/Domain/Entities"
)

type ICategoryRepository interface {
	New()
	CreateCategory(category *entities.Category) *entities.Category
	FindAll() []entities.Category
	FindById(id int) *entities.Category
	DeleteCategory(id int) *entities.Category
	UpdateCategory(id int, attributes map[string]interface{}) *entities.Category
}
