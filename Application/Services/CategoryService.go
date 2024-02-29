package services

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	mysqlrepository "delivery/golang_salesInvoice/Infrastructure/Repositories/MysqlRepository"
)

type CategoryService struct {
	repository mysqlrepository.CategoryRepository
}

func (service CategoryService) getAll() []models.Category {
	var categories []models.Category

}
