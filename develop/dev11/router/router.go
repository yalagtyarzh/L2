package router

import (
	"net/http"

	"dev11/handlers"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", handlers.CreateEvent)
	mux.HandleFunc("/update_event", handlers.UpdateEvent)
	mux.HandleFunc("/delete_event", handlers.DeleteEvent)
	mux.HandleFunc("/events_for_day", handlers.EventsForDay)
	mux.HandleFunc("/events_for_week", handlers.EventsForWeek)
	mux.HandleFunc("/events_for_month", handlers.EventsForMonth)

	return mux
}
