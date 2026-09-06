package handlers

import (
	"html/template"

	webtemplates "github.com/falsisdev/website/web/templates"
)

var tmpl *template.Template

func InitTemplates() error {
	var err error
	tmpl, err = template.New("").ParseFS(
		webtemplates.FS,
		"*.html",
		"layouts/*.html",
		"components/*.html",
	)
	return err
}
