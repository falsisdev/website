package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/falsisdev/website/web"
)

var tmpl *template.Template

var funcMap = template.FuncMap{
	"age": func(milliseconds int64) string {
		pastTime := time.UnixMilli(milliseconds)
		years := time.Since(pastTime).Hours() / (24 * 365.25)
		return fmt.Sprintf("%.2f", years)
	},
}

func init() {
	var err error
	tmpl, err = template.New("").Funcs(funcMap).ParseFS(web.Files,
		"templates/*.html",
		"templates/layouts/*.html",
		"templates/components/*.html",
	)
	if err != nil {
		panic(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) >= 8 && r.URL.Path[:8] == "/static/" {
		http.FileServer(http.FS(web.Files)).ServeHTTP(w, r)
		return
	}

	data := map[string]interface{}{
		"TargetDate": int64(1199002814000),
	}

	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}