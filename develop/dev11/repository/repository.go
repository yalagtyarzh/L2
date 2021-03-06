package repository

import (
	"time"

	"dev11/models"
)

// DatabaseRepo представляет собой интерфейс, диктующий необходимые методы для подключаемых баз данных
type DatabaseRepo interface {
	InsertEvent(e models.Event) error
	UpdateEvent(e models.Event) error
	DeleteEvent(id int) error
	GetEventsForDay(userID int, date time.Time) ([]models.Event, error)
	GetEventsForWeek(userID int, date time.Time) ([]models.Event, error)
	GetEventsForMonth(userID int, date time.Time) ([]models.Event, error)
}
