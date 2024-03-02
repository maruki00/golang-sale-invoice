package mysqlrepository

import (
	entities "saleinvoice/App/Domain/Entities"
	"saleinvoice/App/Infrastructure/database"

	"errors"
)

var (
	errFind   = errors.New("Could Not find Categories")
	errStore  = errors.New("Could Not Create New Category")
	errUpdate = errors.New("Could Not Update Category")
)

type CategoryRepository struct {
	dbHandler *database.DBHandler
}

func (cp *CategoryRepository) New() *CategoryRepository {
	return &CategoryRepository{
		dbHandler: database.NewDB(),
	}
}

func (cp CategoryRepository) CreateCategory(category *entities.Category) (*entities.Category, error) {
	err := cp.dbHandler.DB.Create(category)
	if err != nil {
		return nil, errStore
	}
	return category, nil
}

func (cp CategoryRepository) FindAll() ([]entities.Category, error) {
	var categories []entities.Category
	err := cp.dbHandler.DB.Find(&categories)
	if err != nil {
		return nil, errFind
	}
	return categories, nil
}

func (cp CategoryRepository) FindById(id int) (*entities.Category, error) {
	var category entities.Category
	err := cp.dbHandler.DB.Where("Id = ?", id).Find(&category)
	if err != nil {
		return nil, errFind
	}
	return &category, nil
}

func (cp CategoryRepository) DeleteCategory(id int) (*entities.Category, error) {
	var category entities.Category
	err := cp.dbHandler.DB.Where("Id = ?", id).Delete(&category)
	if err != nil {
		return nil, errFind
	}
	return &category, nil
}

func (cp CategoryRepository) UpdateCategory(id int, attributes map[string]interface{}) (*entities.Category, error) {
	var category entities.Category
	err := cp.dbHandler.DB.Where("Id  = ? ", id).Find(&category)
	if err != nil {
		return nil, errFind
	}
	err = cp.dbHandler.DB.Model(&category).Updates(attributes)
	if err != nil {
		return nil, errUpdate
	}
	return &category, nil
}
