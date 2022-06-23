package dbrepo

import (
	"sync"

	"dev11/models"
)

type MemoryStorage struct {
	identifier int
	mutex      *sync.Mutex
	events     map[int]models.Event
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		identifier: 1,
		mutex:      &sync.Mutex{},
		events:     make(map[int]models.Event),
	}
}
