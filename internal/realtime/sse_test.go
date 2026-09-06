package realtime

import (
	"bufio"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSSEHandlerStreamsEvent(t *testing.T) {
	bus := NewBus()
	server := httptest.NewServer(NewSSEHandler(bus))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("SSE request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	reader := bufio.NewReader(res.Body)
	connected, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("failed to read SSE connection event: %v", err)
	}
	if !strings.Contains(connected, ": connected") {
		t.Fatalf("expected connection comment, got %q", connected)
	}
	if _, err := reader.ReadString('\n'); err != nil {
		t.Fatalf("failed to read connection event separator: %v", err)
	}

	bus.Publish(Event{Type: "github_profile_updated", Data: map[string]any{"name": "falsisdev"}})
	eventLine, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("failed to read SSE event type: %v", err)
	}
	dataLine, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("failed to read SSE event payload: %v", err)
	}
	cancel()

	if !strings.Contains(eventLine, "event: github_profile_updated") {
		t.Fatalf("expected SSE event type, got %q", eventLine)
	}
	if !strings.Contains(dataLine, "falsisdev") {
		t.Fatalf("expected SSE event payload, got %q", dataLine)
	}
}
