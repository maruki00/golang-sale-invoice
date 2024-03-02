package controllers

import (
	models "saleinvoice/Domain/Entities"
	mysqlrepository "saleinvoice/Infrastructure/Repositories/MysqlRepository"
)

type CategoryController struct {
	repository mysqlrepository.CategoryRepository
}

func (c *CategoryController) Index() {

}

func (con CategoryController) index() []models.Category {

}
