package interactors

import (
	contracts "saleinvoice/App/Domain/Contracts"
	entities "saleinvoice/App/Domain/Entities"
	factories "saleinvoice/App/Infrastructure/Factories"
)

type CategoryService struct {
	repository contracts.ICategoryRepository
}

func NewCategoryService(repo contracts.ICategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repo,
	}
}

func (service *CategoryService) Index() []entities.Category {
	castegories := service.repository.FindAll()
	return castegories
}

func (service *CategoryService) Store(attributes map[string]interface{}) models.Category {
	category := factories.CategoryFactory(attributes)
	service.repository.CreateCategory(category)
}
