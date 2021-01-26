package app

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Code            int    `json:"code,omitempty"`
	Message         string `json:"message,omitempty"`
	InternalMessage string `json:"internal_message,omitempty"`
}

func RenderJSON(w http.ResponseWriter, any interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(any)
}

func RenderJSONwithStatus(w http.ResponseWriter, status int, any interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(any)
}

func RenderErrorJSON(w http.ResponseWriter, err *Error) {
	var e apiError
	e.Message = err.message
	e.InternalMessage = err.err.Error()
	e.Code = err.code
	if err.code == 0 {
		e.Code = http.StatusBadRequest
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	json.NewEncoder(w).Encode(e)
}
