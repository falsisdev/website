package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/falsisdev/website/handlers"
)

func main() {
	if err := handlers.InitTemplates(); err != nil {
		panic(fmt.Sprintf("An error occured while parsing template files: %v", err))
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("GET /", handlers.HomeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running at http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}

// Local development:
// 1. ./build.sh
// 2. .bin/tailwindcss -i web/static/css/input.css -o web/static/css/output.css --watch
// 3. go run ./cmd/server
