package utils

import (
	"encoding/json"
	"net/http"
)

func ReadBody[T any](r *http.Request) (T, error) {
	var Body T
	err := json.NewDecoder(r.Body).Decode(&Body)
	return Body, err
}
