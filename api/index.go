package handler

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"time"
)

var files embed.FS

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
	tmpl, err = template.New("").Funcs(funcMap).ParseFS(files,
		"web/templates/*.html",
		"web/templates/layouts/*.html",
		"web/templates/components/*.html",
	)
	if err != nil {
		panic(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) >= 8 && r.URL.Path[:8] == "/static/" {
		staticSubFS, err := fs.Sub(files, "web")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.FileServer(http.FS(staticSubFS)).ServeHTTP(w, r)
		return
	}

	data := map[string]any{
		"TargetDate": int64(1199002814000),
	}

	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
