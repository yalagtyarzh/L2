package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"dev11/helpers"
	"dev11/models"
	"dev11/repository"
	"dev11/validation"
)

// Repository представляет собой класс приложения
type Repository struct {
	DB repository.DatabaseRepo
}

// Repo для обработки запросов с помощью встроенных в него структур
var Repo *Repository

// NewRepo возвращает новый Repository
func NewRepo(repo repository.DatabaseRepo) *Repository {
	return &Repository{DB: repo}
}

// NewHandler инициализирует переданный Repository в обработчиках
func NewHandler(r *Repository) {
	Repo = r
}

// CreateEvent обрабатывает запрос на создание события
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		helpers.ThrowError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	event, err := helpers.UnmarshalEvent(r)
	if err != nil {
		helpers.ThrowError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal event: %s", err))
		return
	}

	err = Repo.DB.InsertEvent(event)
	if err != nil {
		helpers.ThrowError(w, http.StatusGatewayTimeout, fmt.Errorf("insert event error: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	helpers.WriteResponse(w, http.StatusOK, "Event created successfully!", []models.Event{event})
}

// UpdateEvent обрабатывает запрос на обновление события
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		helpers.ThrowError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	event, err := helpers.UnmarshalEvent(r)
	if err != nil {
		helpers.ThrowError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal event: %s", err))
		return
	}

	err = Repo.DB.UpdateEvent(event)
	if err != nil {
		helpers.ThrowError(w, http.StatusGatewayTimeout, fmt.Errorf("update event error: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	helpers.WriteResponse(w, http.StatusOK, "Event updated successfully!", []models.Event{event})
}

// DeleteEvent обрабатывает запрос на удаление события
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		helpers.ThrowError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	event, err := helpers.UnmarshalEvent(r)
	if err != nil {
		helpers.ThrowError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal event: %s", err))
		return
	}

	err = Repo.DB.DeleteEvent(event.ID)
	if err != nil {
		helpers.ThrowError(w, http.StatusGatewayTimeout, fmt.Errorf("insert event error: %s", err))
		return
	}

	helpers.WriteResponse(w, http.StatusOK, "Event deleted successfully!", []models.Event{event})
}

// dateLayout для парсинга времени, переданного в Get запросах
const dateLayout = "2006-01-02"

// EventsForDay обрабатывает запрос на получение событий за день
func EventsForDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		helpers.ThrowError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	required := []string{"user_id", "date"}

	v := r.URL.Query()
	values, err := validation.ValidateURLValues(v, required...)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	date, err := time.Parse(dateLayout, values["date"])
	if err != nil {
		helpers.ThrowError(
			w, http.StatusBadRequest, fmt.Errorf("invalid date format: %s, waited for %s", values["date"], dateLayout),
		)
		return
	}

	userID, err := strconv.Atoi(values["user_id"])
	if err != nil || userID < 1 {
		helpers.ThrowError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %s", values["user_id"]))
		return
	}

	events, err := Repo.DB.GetEventsForDay(userID, date)
	if err != nil {
		helpers.ThrowError(w, http.StatusGatewayTimeout, fmt.Errorf("get events for day error: %s", err))
		return
	}

	helpers.WriteResponse(w, http.StatusOK, "Got events!", events)
}

// EventsForWeek обрабатывает запрос на получение событий за неделю
func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		helpers.ThrowError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	required := []string{"user_id", "date"}

	v := r.URL.Query()
	values, err := validation.ValidateURLValues(v, required...)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	date, err := time.Parse(dateLayout, values["date"])
	if err != nil {
		helpers.ThrowError(
			w, http.StatusBadRequest, fmt.Errorf("invalid date format: %s, waited for %s", values["date"], dateLayout),
		)
		return
	}

	userID, err := strconv.Atoi(values["user_id"])
	if err != nil || userID < 1 {
		helpers.ThrowError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %s", values["user_id"]))
		return
	}

	events, err := Repo.DB.GetEventsForWeek(userID, date)
	if err != nil {
		helpers.ThrowError(w, http.StatusGatewayTimeout, fmt.Errorf("get events for day error: %s", err))
		return
	}

	helpers.WriteResponse(w, http.StatusOK, "Got events!", events)
}

// EventsForMonth обрабатывает запрос на получение событий за месяц
func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		helpers.ThrowError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	required := []string{"user_id", "date"}

	v := r.URL.Query()
	values, err := validation.ValidateURLValues(v, required...)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	date, err := time.Parse(dateLayout, values["date"])
	if err != nil {
		helpers.ThrowError(
			w, http.StatusBadRequest, fmt.Errorf("invalid date format: %s, waited for %s", values["date"], dateLayout),
		)
		return
	}

	userID, err := strconv.Atoi(values["user_id"])
	if err != nil || userID < 1 {
		helpers.ThrowError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %s", values["user_id"]))
		return
	}

	events, err := Repo.DB.GetEventsForMonth(userID, date)
	if err != nil {
		helpers.ThrowError(w, http.StatusGatewayTimeout, fmt.Errorf("get events for day error: %s", err))
		return
	}

	helpers.WriteResponse(w, http.StatusOK, "Got events!", events)
}
