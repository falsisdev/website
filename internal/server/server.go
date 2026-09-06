package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/falsisdev/website/handlers"
	"github.com/falsisdev/website/internal/config"
)

func NewMux(cfg config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", HealthHandler)

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("GET /", handlers.HomeHandler)

	_ = cfg
	return mux
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	payload := map[string]any{
		"status":      "ok",
		"service":     "website",
		"timestamp":   time.Now().UTC().Format(time.RFC3339),
		"environment": "development",
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "failed to encode health payload", http.StatusInternalServerError)
	}
}
