package realtime

import "testing"

func TestBusPublishesToSubscribers(t *testing.T) {
	bus := NewBus()
	sub := bus.Subscribe()
	defer bus.Unsubscribe(sub)

	bus.Publish(Event{Type: "github_activity", Data: map[string]string{"status": "ok"}})

	select {
	case event := <-sub:
		if event.Type != "github_activity" {
			t.Fatalf("expected github_activity event, got %q", event.Type)
		}
	default:
		t.Fatal("expected event to be published to subscriber")
	}
}
