package helpers

import (
	"encoding/json"
	"net/http"

	"dev11/models"
)

// responseOK представляет собой класс валидного json ответа
type responseOK struct {
	Message string         `json:"message"`
	Events  []models.Event `json:"events"`
}

// responseError представляет собой класс невалидного json ответа
type responseError struct {
	Error string `json:"error"`
}

// ThrowError прокидывает в ответ responseError с переданным статусом ответа и ошибкой
func ThrowError(w http.ResponseWriter, status int, err error) {
	resp := responseError{
		Error: err.Error(),
	}

	out, _ := json.MarshalIndent(resp, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
}

// WriteResponse прокидывает ответ responseOK с переданным статусом, сообщением и срезом событий
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
