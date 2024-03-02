package mysqlrepository

import (
	models "saleinvoice/Domain/Entities"
	"saleinvoice/Infrastructure/database"

	"errors"

	"gorm.io/gorm"
)

var (
	errFind   = errors.New("Could Not find Categories")
	errStore  = errors.New("Could Not Create New Category")
	errUpdate = errors.New("Could Not Update Category")
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

func (cp CategoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	err := cp.db.Create(category)
	if err != nil {
		return nil, errStore
	}
	return category, nil
}

func (cp CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := cp.db.Find(&categories)
	if err != nil {
		return nil, errFind
	}
	return categories, nil
}

func (cp CategoryRepository) FindById(id int) (*models.Category, error) {
	var category models.Category
	err := cp.db.Where("Id = ?", id).Find(&category)
	if err != nil {
		return nil, errFind
	}
	return &category, nil
}

func (cp CategoryRepository) DeleteCategory(id int) (*models.Category, error) {
	var category models.Category
	err := cp.db.Where("Id = ?", id).Delete(&category)
	if err != nil {
		return nil, errFind
	}
	return &category, nil
}

func (cp CategoryRepository) UpdateCategory(id int, attributes map[string]interface{}) (*models.Category, error) {
	var category models.Category
	err := cp.db.Where("Id  = ? ", id).Find(&category)
	if err != nil {
		return nil, errFind
	}
	err = cp.db.Model(&category).Updates(attributes)
	if err != nil {
		return nil, errUpdate
	}
	return &category, nil
}
