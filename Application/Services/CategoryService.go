package services

import (
	contracts "saleinvoice/Domain/Contracts"
	models "saleinvoice/Domain/Entities"
	mysqlrepository "saleinvoice/Infrastructure/Repositories/MysqlRepository"
)

type CategoryService struct {
	repository mysqlrepository.CategoryRepository
}

func NewCategoryService(repository contracts.ICategoryRepository)
func (service CategoryService) getAll() []models.Category {
	var categories []models.Category
	service.repository.FindAll(&categories)
	return categories
}
