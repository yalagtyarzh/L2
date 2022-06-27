package dbrepo

import (
	"sync"

	"dev11/models"
)

// MemoryStorage представляет собой хранилище данных в памяти
type MemoryStorage struct {
	identifier int
	mutex      *sync.Mutex
	events     map[int]models.Event
}

// NewMemoryStorage возвращает новый MemoryStorage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		identifier: 1,
		mutex:      &sync.Mutex{},
		events:     make(map[int]models.Event),
	}
}
