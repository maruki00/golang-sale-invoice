package services

import (
	contracts "saleinvoice/Domain/Contracts"
	models "saleinvoice/Domain/Entities"
	factories "saleinvoice/Infrastructure/Factories"
)

type CategoryService struct {
	repository contracts.ICategoryRepository
}

func NewCategoryService(repo contracts.ICategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repo,
	}
}

func (service *CategoryService) Index() []models.Category {
	castegories := service.repository.FindAll()
	return castegories
}

func (service *CategoryService) Store(attributes map[string]interface{}) models.Category {
	category := factories.CategoryFactory(attributes)
}
