package cache

import (
	"sync"
	"time"
)

type Item[T any] struct {
	Value     T
	ExpiresAt time.Time
	FetchedAt time.Time
}

type Store[T any] struct {
	mu    sync.RWMutex
	items map[string]Item[T]
}

func NewStore[T any]() *Store[T] {
	return &Store[T]{items: make(map[string]Item[T])}
}

func (s *Store[T]) Get(key string) (T, bool) {
	s.mu.RLock()
	item, ok := s.items[key]
	s.mu.RUnlock()
	if !ok {
		var zero T
		return zero, false
	}
	if time.Now().After(item.ExpiresAt) {
		s.mu.Lock()
		delete(s.items, key)
		s.mu.Unlock()
		var zero T
		return zero, false
	}
	return item.Value, true
}

func (s *Store[T]) Set(key string, value T, ttl time.Duration) {
	s.mu.Lock()
	s.items[key] = Item[T]{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
		FetchedAt: time.Now(),
	}
	s.mu.Unlock()
}

func (s *Store[T]) Delete(key string) {
	s.mu.Lock()
	delete(s.items, key)
	s.mu.Unlock()
}
