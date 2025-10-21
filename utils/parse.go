package utils

import (
	"encoding/json"
)

func parseRequest(r *http.Request, target interface{}) error {
	defer r.Body.Close()
	return json.NewDecode(r.Body).Decode(target)
}