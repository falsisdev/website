package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/falsisdev/website/handlers"
	"github.com/falsisdev/website/internal/config"
	"github.com/falsisdev/website/internal/realtime"
)

func NewMux(cfg config.Config) http.Handler {
	return NewMuxWithBus(cfg, realtime.NewBus())
}

func NewMuxWithBus(cfg config.Config, bus *realtime.Bus) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", HealthHandler)
	mux.HandleFunc("GET /healthz", HealthHandler)
	mux.HandleFunc("GET /events/github", realtime.NewSSEHandler(bus))

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandlerWithConfig(cfg, w, r)
	})

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
