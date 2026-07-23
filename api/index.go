package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/falsisdev/website/web"
)

var (
	tmpl    *template.Template
	initErr error
)

var funcMap = template.FuncMap{
	"age": func(milliseconds int64) string {
		pastTime := time.UnixMilli(milliseconds)
		years := time.Since(pastTime).Hours() / (24 * 365.25)
		return fmt.Sprintf("%.2f", years)
	},
}

func init() {
	tmpl, initErr = template.New("").Funcs(funcMap).ParseFS(web.Files,
		"templates/*.html",
		"templates/layouts/*.html",
		"templates/components/*.html",
	)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if initErr != nil {
		http.Error(w, fmt.Sprintf("Template Init Error: %v", initErr), http.StatusInternalServerError)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/static/") {
		http.FileServer(http.FS(web.Files)).ServeHTTP(w, r)
		return
	}

	data := map[string]interface{}{
		"TargetDate": int64(1199002814000),
	}

	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Render Error: %v", err), http.StatusInternalServerError)
	}
}