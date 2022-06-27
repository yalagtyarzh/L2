package dbrepo

import (
	"errors"
	"time"

	"dev11/models"
)

// InsertEvent вставляет event в Storage
func (m *MemoryStorage) InsertEvent(e models.Event) error {
	m.mutex.Lock()
	e.ID = m.identifier
	m.events[e.ID] = e
	m.identifier++
	m.mutex.Unlock()
	return nil
}

// UpdateEvent обновляет event в Storage
func (m *MemoryStorage) UpdateEvent(e models.Event) error {
	_, ok := m.events[e.ID]
	if !ok {
		return errors.New("event not found")
	}

	m.events[e.ID] = e
	return nil
}

// DeleteEvent удаляет event в Storage по переданному id
func (m *MemoryStorage) DeleteEvent(id int) error {
	_, ok := m.events[id]
	if !ok {
		return errors.New("event not found")
	}

	delete(m.events, id)
	return nil
}

// GetEventsForDay возвращает все события для определенного пользователя в определенном времени в разрезе дня
func (m *MemoryStorage) GetEventsForDay(userID int, date time.Time) ([]models.Event, error) {
	events := make([]models.Event, 0)
	m.mutex.Lock()

	for _, event := range m.events {
		if event.Date.Equal(date) && userID == event.UserID {
			events = append(events, event)
		}
	}
	m.mutex.Unlock()
	return events, nil
}

// GetEventsForWeek возвращает все события для определенного пользователя в определенном времени в разрезе недели
func (m *MemoryStorage) GetEventsForWeek(userID int, date time.Time) ([]models.Event, error) {
	events := make([]models.Event, 0)
	m.mutex.Lock()

	for _, event := range m.events {
		EventY, EventW := event.Date.ISOWeek()
		DateY, DateW := date.ISOWeek()
		if EventY == DateY && EventW == DateW && userID == event.UserID {
			events = append(events, event)
		}
	}

	m.mutex.Unlock()
	return events, nil
}

// GetEventsForMonth возвращает все события для определенного пользователя в определенном времени в разрезе месяца
func (m *MemoryStorage) GetEventsForMonth(userID int, date time.Time) ([]models.Event, error) {
	events := make([]models.Event, 0)
	m.mutex.Lock()

	for _, event := range m.events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.UserID == userID {
			events = append(events, event)
		}
	}
	m.mutex.Unlock()
	return events, nil
}
