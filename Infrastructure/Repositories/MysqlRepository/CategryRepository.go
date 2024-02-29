package mysqlrepository

import (
	models "saleinvoice/Domain/Entities"
	"saleinvoice/Infrastructure/database"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func (cp *CategoryRepository) New() *CategoryRepository {
	if cp.db == nil {
		cp.db = database.Connect()
	}
	return cp
}

func (cp CategoryRepository) CreateCategory(category *models.Category) *models.Category {
	cp.db.Create(category)
	return category
}

func (cp CategoryRepository) FindAll() []models.Category {
	var categories []models.Category
	cp.db.Find(&categories)
	return categories
}

func (cp CategoryRepository) FindById(id int) *models.Category {
	var category models.Category
	cp.db.Where("Id = ?", id).Find(&category)
	return &category
}

func (cp CategoryRepository) DeleteCategory(id int) *models.Category {
	var category models.Category
	cp.db.Where("Id = ?", id).Delete(&category)
	return &category
}

func (cp CategoryRepository) UpdateCategory(id int, attributes map[string]interface{}) *models.Category {
	var category models.Category
	cp.db.Where("Id  = ? ", id).Find(&category)
	cp.db.Model(&category).Updates(attributes)
	return &category
}
