package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPageDataEnglish(t *testing.T) {
	data := pageData("en")
	if data.Language != "en" {
		t.Fatalf("expected English language, got %q", data.Language)
	}
	if len(data.Projects) == 0 {
		t.Fatal("expected English project list to be populated")
	}
	if data.HeroTitle == "" {
		t.Fatal("expected English hero title to be populated")
	}
}

func TestPageDataTurkish(t *testing.T) {
	data := pageData("tr")
	if data.Language != "tr" {
		t.Fatalf("expected Turkish language, got %q", data.Language)
	}
	if data.LanguageLabel != "Dil" {
		t.Fatalf("expected Turkish language label, got %q", data.LanguageLabel)
	}
}

func TestLanguageFromRequestSetsCookieAndDefaultLanguage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?lang=tr", nil)
	res := httptest.NewRecorder()

	language := languageFromRequest(res, req)
	if language != "tr" {
		t.Fatalf("expected language tr, got %q", language)
	}
	if cookie := res.Result().Cookies(); len(cookie) == 0 {
		t.Fatal("expected language cookie to be set")
	}
}
