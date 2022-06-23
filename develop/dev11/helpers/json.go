package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"dev11/models"
)

func UnmarshalEvent(r *http.Request) (models.Event, error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return models.Event{}, err
	}

	var event models.Event

	err = json.Unmarshal(b, &event)
	if err != nil {
		return models.Event{}, err
	}

	if event.ID < 1 || event.UserID < 1 {
		return models.Event{}, fmt.Errorf("invalid event ID or user ID")
	}

	return event, nil
}
