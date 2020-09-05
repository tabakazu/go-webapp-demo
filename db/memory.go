package db

import (
	"context"
	"sync"

	"github.com/tabakazu/golang-webapi-demo/model"
)

type memoryDB struct {
	db   map[string]*model.Item
	lock sync.RWMutex
}

func NewMemoryDB() *memoryDB {
	return &memoryDB{db: map[string]*model.Item{}}
}

func (m *memoryDB) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result := make([]*model.Item, len(m.db))
	i := 0
	for _, t := range m.db {
		result[i] = t
		i++
	}

	return result, nil
}

func (m *memoryDB) PutItem(ctx context.Context, t *model.Item) error {
	m.lock.Lock()
	m.db[t.ID] = t
	m.lock.Unlock()
	return nil
}
