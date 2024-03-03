package interactors

import (
	"fmt"
	"log"
	"net/http"
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

func (service *CategoryInteractor) Store(request *http.Request) *entities.Category {
	fmt.Println("iNTERCATOR BODY: ", request.Body)
	category := factories.CategoryFactory(request)
	fmt.Println("interasctor: ", category)
	service.repository.CreateCategory(category)
	return category
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
