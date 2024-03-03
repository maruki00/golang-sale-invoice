package factories

import (
	"fmt"
	"io"
	"log"
	"net/http"
	ej "random/golang/easyJson/model"
	entities "saleinvoice/App/Domain/Entities"

	"github.com/mailru/easyjson"
)

func CategoryFactory(request *http.Request) *entities.Category {

	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal("Error Parse Body")
	}
	err = ej.MarshalerUnmarshaler([]byte(body))
	easyjson.Unmarshal([]byte(body))
	fmt.Println("factory: ", category)
	return category
}
