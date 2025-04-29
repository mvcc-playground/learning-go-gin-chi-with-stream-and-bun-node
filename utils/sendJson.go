package utils

import (
	"encoding/json"
	"net/http"
)

func SendJson[Data any](w http.ResponseWriter, status int, data Data) {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
