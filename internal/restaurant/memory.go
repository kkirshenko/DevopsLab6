package restaurant

import (
	"context"
	"sync"
)

type MemoryRepository struct {
	mu sync.RWMutex
	items map[string]Restaurant
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{items: make(map[string]Restaurant)}
}

func (r *MemoryRepository) Create(_ context.Context, item Restaurant) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if item.ID == "" { return "", ErrInvalid }
	if _, exists := r.items[item.ID]; exists { return "", ErrInvalid }
	r.items[item.ID] = item
	return item.ID, nil
}

func (r *MemoryRepository) GetByID(_ context.Context, id string) (Restaurant, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[id]
	if !ok { return Restaurant{}, ErrNotFound }
	return item, nil
}

func (r *MemoryRepository) ChangeRating(_ context.Context, id string, rating float32) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	item, ok := r.items[id]
	if !ok { return ErrNotFound }
	item.Rating = rating
	r.items[id] = item
	return nil
}
