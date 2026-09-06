package main

import (
	"fmt"
	"net/http"

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

	fmt.Println("Server running at the port :8080... http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

// Local development:
// 1. ./build.sh
// 2. .bin/tailwindcss -i web/static/css/input.css -o web/static/css/output.css --watch
// 3. go run ./cmd/server
