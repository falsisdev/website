package sanity

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListPosts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST method, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"result":[{"_id":"post-1","title":"Hello Sanity","summary":"Summary here","slug":{"current":"hello-sanity"},"publishedAt":"2026-01-01T00:00:00Z"}]}`))
	}))
	defer server.Close()

	client := New(Config{
		ProjectID:  "demo-project",
		Dataset:    "production",
		APIVersion: "2024-01-01",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	})

	posts, err := client.ListPosts(context.Background())
	if err != nil {
		t.Fatalf("ListPosts() returned unexpected error: %v", err)
	}
	if len(posts) != 1 {
		t.Fatalf("expected 1 post, got %d", len(posts))
	}
	if posts[0].Slug != "hello-sanity" {
		t.Fatalf("expected slug hello-sanity, got %q", posts[0].Slug)
	}
	if posts[0].Title != "Hello Sanity" {
		t.Fatalf("expected title Hello Sanity, got %q", posts[0].Title)
	}
}
