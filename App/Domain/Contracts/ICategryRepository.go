package contracts

import (
	entities "saleinvoice/App/Domain/Entities"
)

type ICategoryRepository interface {
	CreateCategory(category *entities.Category) (*entities.Category, error)
	FindAll() ([]entities.Category, error)
	FindById(id int) (*entities.Category, error)
	DeleteCategory(id int) (*entities.Category, error)
	UpdateCategory(id int, attributes map[string]interface{}) (*entities.Category, error)
}
