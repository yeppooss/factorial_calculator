package pr

import (
	"encoding/json"
	"factorial_calculate/internal/http-server/requests"
	"net/http"
)

func ParseRequest(r *http.Request) (requests.RequestBody, error) {
	var reqBody requests.RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		return requests.RequestBody{}, err
	}

	return reqBody, nil

}
