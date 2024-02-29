package contracts

import models "saleinvoice/Domain/Entities"

type ICategoryRepository interface {
	New()
	CreateCategory(category models.Category) *models.Category
	FindAll() []models.Category
	FindById(id int) *models.Category
	DeleteCategory(id int) *models.Category
	UpdateCategory(id int, attributes map[string]interface{}) *models.Category
}
