package github

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/falsisdev/website/internal/realtime"
)

func TestPollerPublishesProfileUpdate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/falsisdev" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"login":"falsisdev","name":"Falsis","avatar_url":"https://example.com/avatar.png","bio":"Backend engineer","followers":12,"following":8,"public_repos":5}`))
	}))
	defer server.Close()

	bus := realtime.NewBus()
	sub := bus.Subscribe()
	defer bus.Unsubscribe(sub)

	poller := NewPoller(PollerConfig{
		Username:   "falsisdev",
		Token:      "",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
		Interval:   250 * time.Millisecond,
		Bus:        bus,
	})

	if err := poller.PollOnce(context.Background()); err != nil {
		t.Fatalf("PollOnce() returned error: %v", err)
	}

	select {
	case event := <-sub:
		if event.Type != "github_profile_updated" {
			t.Fatalf("expected github_profile_updated event, got %q", event.Type)
		}
		data, ok := event.Data.(map[string]any)
		if !ok {
			t.Fatalf("expected map payload, got %T", event.Data)
		}
		if data["name"] != "Falsis" {
			t.Fatalf("expected profile name Falsis, got %v", data["name"])
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for github_profile_updated event")
	}
}
