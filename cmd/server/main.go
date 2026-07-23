package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go Sunucusu Çalışıyor!")
	})

	fmt.Println("Sunucu 8080 portunda başlatılıyor... http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// Vercel Deploy için build command kısmına yazılacak: tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --minify && go build -o server cmd/server/main.go
//Geiştirme anında bir terminal sekmesine yazılacak ve terminal sekmesi aktif tutulacak: tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --watch