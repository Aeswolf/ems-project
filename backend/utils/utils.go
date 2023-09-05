package utils

import (
	"encoding/json"
	"net/http"
)

type apiHandler func(http.ResponseWriter, *http.Request) error

type errorHandler struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

func MakeHTTPHandler(f apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, errorHandler{Error: err.Error()})
			return
		}
	}
}
