package handlers

import (
	"fmt"
	"html/template"
	"path/filepath"
	"time"
)

var tmpl *template.Template

var funcMap = template.FuncMap{
	"age": func(milliseconds int64) string {
		pastTime := time.UnixMilli(milliseconds)
		years := time.Since(pastTime).Hours() / (24 * 365.25)
		return fmt.Sprintf("%.2f", years)
	},
}

func InitTemplates() error {
	var files []string

	patterns := []string{
		"web/templates/*.html",
		"web/templates/layouts/*.html",
		"web/templates/components/*.html",
	}

	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return err
		}
		files = append(files, matches...)
	}

	if len(files) == 0 {
		return fmt.Errorf("hiçbir HTML şablon dosyası bulunamadı")
	}

	var err error
	tmpl, err = template.New("").Funcs(funcMap).ParseFiles(files...)
	return err
}