package utils

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, status int, Response interface{}) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(Response)
}
