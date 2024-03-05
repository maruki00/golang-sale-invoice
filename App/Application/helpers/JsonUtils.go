package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ParseBody(request *http.Request, x interface{}) {
	if body, err := io.ReadAll(request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			log.Fatal(err.Error())
			fmt.Println("Utils Broken: ", x)
			return
		}
	}

}
