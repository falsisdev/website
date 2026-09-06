package handler

import (
	"net/http"
	"sync"

	"github.com/falsisdev/website/internal/handlers"
)

var (
	templateOnce sync.Once
	templateErr  error
)

// Handler is the Vercel entrypoint for the Go serverless function.
func Handler(w http.ResponseWriter, r *http.Request) {
	templateOnce.Do(func() {
		templateErr = handlers.InitTemplates()
	})
	if templateErr != nil {
		http.Error(w, templateErr.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	handlers.HomeHandler(w, r)
}
