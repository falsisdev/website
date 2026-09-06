package github

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListRepos(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/falsisdev/repos" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[
			{"name":"website","description":"portfolio","html_url":"https://github.com/falsisdev/website","language":"Go"},
			{"name":"notes","description":"notes","html_url":"https://github.com/falsisdev/notes","language":"Markdown"}
		]`))
	}))
	defer server.Close()

	client := New(Config{
		Username:   "falsisdev",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	})

	repos, err := client.ListRepos(context.Background())
	if err != nil {
		t.Fatalf("ListRepos() returned unexpected error: %v", err)
	}
	if len(repos) != 2 {
		t.Fatalf("expected 2 repos, got %d", len(repos))
	}
	if repos[0].Name != "website" {
		t.Fatalf("expected first repo name website, got %q", repos[0].Name)
	}
}
