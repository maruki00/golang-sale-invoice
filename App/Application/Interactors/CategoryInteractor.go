package interactors

import (
	"log"
	entities "saleinvoice/App/Domain/Entities"
	factories "saleinvoice/App/Infrastructure/Factories"
	mysqlrepository "saleinvoice/App/Infrastructure/Repositories/MysqlRepository"
)

type CategoryInteractor struct {
	repository *mysqlrepository.CategoryRepository
}

func NewCatergoryInteractor(repository *mysqlrepository.CategoryRepository) *CategoryInteractor {
	return &CategoryInteractor{
		repository: repository,
	}
}

func (interactor *CategoryInteractor) Index() []entities.Category {
	categoories, err := interactor.repository.FindAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	return categoories
}

func (service *CategoryInteractor) Store(attributes map[string]interface{}) entities.Category {
	category := factories.CategoryFactory(attributes)
	service.repository.CreateCategory(category)
	return *category
}

func (interactor CategoryInteractor) GetByID(id int) *entities.Category {
	category, err := interactor.repository.FindById(id)
	if err != nil {
		log.Fatal("Category not found")
	}
	return category
}

func (interactor CategoryInteractor) Delete(id int) *entities.Category {
	category, err := interactor.repository.DeleteCategory(id)
	if err != nil {
		log.Fatal("Category not found")
	}
	return category
}
