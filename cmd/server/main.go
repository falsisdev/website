package main

import (
	"fmt"
	"net/http"

	"github.com/falsisdev/website/internal/handlers"
)

func main() {
	if err := handlers.InitTemplates(); err != nil {
		panic(fmt.Sprintf("Şablonlar yüklenirken hata oluştu: %v", err))
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("GET /", handlers.HomeHandler)

	fmt.Println("Sunucu 8080 portunda çalışıyor... http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

// Vercel Deploy için build command kısmına yazılacak: tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --minify && go build -o server cmd/server/main.go
//Geiştirme anında bir terminal sekmesine yazılacak ve terminal sekmesi aktif tutulacak: tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --watch