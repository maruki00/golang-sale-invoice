package helpers

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(request *http.Request, x interface{}) {
	if body, err := io.ReadAll(request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func RequestToMap(request *http.Request, x *map[string]interface{}) {
	if body, err := io.ReadAll(request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
