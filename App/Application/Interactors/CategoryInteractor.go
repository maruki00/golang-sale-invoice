package interactors

import (
	"log"
	contracts "saleinvoice/App/Domain/Contracts"
	entities "saleinvoice/App/Domain/Entities"
	factories "saleinvoice/App/Infrastructure/Factories"
	mysqlrepository "saleinvoice/App/Infrastructure/Repositories/MysqlRepository"
)

type CategoryInteractor struct {
	repository *mysqlrepository.CategoryRepository
}

func (ci *CategoryInteractor) NewCatergoryInteractor(repository *contracts.ICategoryRepository) *CategoryInteractor {
	return &CategoryInteractor{
		repository: ci.repository,
	}
}

func (interactor *CategoryInteractor) Index() []entities.Category {
	categoories, err := interactor.repository.FindAll()
	if err != nil {
		return nil
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
