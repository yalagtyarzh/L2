package helpers

import (
	"encoding/json"
	"net/http"

	"dev11/models"
)

type responseOK struct {
	Message string         `json:"message"`
	Events  []models.Event `json:"events"`
}

type responseError struct {
	Error string `json:"error"`
}

func ThrowError(w http.ResponseWriter, status int, err error) {
	resp := responseError{
		Error: err.Error(),
	}

	out, _ := json.MarshalIndent(resp, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
}

func WriteResponse(w http.ResponseWriter, status int, msg string, events []models.Event) {
	resp := responseOK{
		Message: msg,
		Events:  events,
	}

	out, _ := json.MarshalIndent(resp, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
}
