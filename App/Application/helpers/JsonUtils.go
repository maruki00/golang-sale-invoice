package helpers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mailru/easyjson"
)

func ParseBody(request *http.Request, x interface{}) {
	if body, err := io.ReadAll(request.Body); err == nil {
		if err := easyjson.Unmarshal([]byte(body), &x); err != nil {
			log.Fatal(err.Error())
			fmt.Println("Utils Broken: ", x)
			return
		}
	}

}
