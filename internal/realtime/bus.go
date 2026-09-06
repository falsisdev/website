package realtime

import (
	"sync"
)

type Event struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type Subscriber chan Event

type Bus struct {
	mu          sync.RWMutex
	subscribers map[Subscriber]struct{}
}

func NewBus() *Bus {
	return &Bus{subscribers: make(map[Subscriber]struct{})}
}

func (b *Bus) Subscribe() Subscriber {
	ch := make(chan Event, 16)
	b.mu.Lock()
	b.subscribers[ch] = struct{}{}
	b.mu.Unlock()
	return ch
}

func (b *Bus) Unsubscribe(ch Subscriber) {
	b.mu.Lock()
	delete(b.subscribers, ch)
	b.mu.Unlock()
	close(ch)
}

func (b *Bus) Publish(event Event) {
	b.mu.RLock()
	subs := make([]Subscriber, 0, len(b.subscribers))
	for sub := range b.subscribers {
		subs = append(subs, sub)
	}
	b.mu.RUnlock()

	for _, sub := range subs {
		select {
		case sub <- event:
		default:
			go b.Unsubscribe(sub)
		}
	}
}
