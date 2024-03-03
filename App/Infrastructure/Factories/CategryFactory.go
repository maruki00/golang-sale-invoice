package factories

import (
	"fmt"
	"net/http"
	"saleinvoice/App/Application/helpers"
	entities "saleinvoice/App/Domain/Entities"
)

func CategoryFactory(request *http.Request) *entities.Category {
	var category entities.Category
	helpers.ParseBody(request, &category)
	fmt.Println("factory: ", category)
	return &category
}
