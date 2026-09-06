package cache

import (
	"testing"
	"time"
)

func TestStoreSetAndGet(t *testing.T) {
	store := NewStore[string]()
	store.Set("greeting", "hello", time.Minute)

	value, ok := store.Get("greeting")
	if !ok {
		t.Fatal("expected greeting to be cached")
	}
	if value != "hello" {
		t.Fatalf("expected hello, got %q", value)
	}
}

func TestStoreExpiresEntries(t *testing.T) {
	store := NewStore[string]()
	store.Set("soon", "value", 10*time.Millisecond)

	time.Sleep(50 * time.Millisecond)
	if _, ok := store.Get("soon"); ok {
		t.Fatal("expected expired cache entry to be removed")
	}
}
