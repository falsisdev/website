package handlers

import (
	"net/http"
)

type PageData struct {
	Language      string
	Title         string
	Brand         string
	Home          string
	LanguageLabel string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	language := languageFromRequest(w, r)
	data := pageData(language)
	templateName := "index.html"
	if r.Header.Get("HX-Request") == "true" {
		templateName = "page"
	}

	err := tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func languageFromRequest(w http.ResponseWriter, r *http.Request) string {
	language := r.URL.Query().Get("lang")
	if language == "" {
		if cookie, err := r.Cookie("language"); err == nil {
			language = cookie.Value
		}
	}
	if language != "tr" {
		language = "en"
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "language",
		Value:    language,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 365,
		SameSite: http.SameSiteLaxMode,
	})
	return language
}

func pageData(language string) PageData {
	if language == "tr" {
		return PageData{
			Language:      "tr",
			Title:         "Kişisel Web Sitesi",
			Brand:         "Kişisel Web Sitesi",
			Home:          "Ana Sayfa",
			LanguageLabel: "Dil",
		}
	}

	return PageData{
		Language:      "en",
		Title:         "Personal Website",
		Brand:         "Personal Website",
		Home:          "Home",
		LanguageLabel: "Language",
	}
}
