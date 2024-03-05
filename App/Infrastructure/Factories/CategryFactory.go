package factories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	entities "saleinvoice/App/Domain/Entities"
)

func CategoryFactory(request *http.Request) *entities.Category {
	body, _ := io.ReadAll(request.Body)
	var category entities.Category
	json.Unmarshal([]byte(body), &category)
	fmt.Println("factory: ", category)
	return &category
}
