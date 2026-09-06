package github

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProfile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/falsisdev" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"login":"falsisdev","name":"Falsis","avatar_url":"https://example.com/avatar.png","bio":"Backend engineer","followers":12,"following":8,"public_repos":5}`))
	}))
	defer server.Close()

	client := New(Config{Username: "falsisdev", BaseURL: server.URL, HTTPClient: server.Client()})
	profile, err := client.GetProfile(context.Background())
	if err != nil {
		t.Fatalf("GetProfile() returned unexpected error: %v", err)
	}
	if profile.Login != "falsisdev" {
		t.Fatalf("expected login falsisdev, got %q", profile.Login)
	}
	if profile.AvatarURL == "" {
		t.Fatal("expected avatar URL to be set")
	}
	if profile.Name != "Falsis" {
		t.Fatalf("expected name Falsis, got %q", profile.Name)
	}
}
