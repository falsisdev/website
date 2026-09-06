package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/falsisdev/website/internal/config"
	"github.com/falsisdev/website/internal/github"
	"github.com/falsisdev/website/internal/sanity"
)

type Project struct {
	Name        string
	Description string
	Stack       string
	Link        string
}

type BlogPost struct {
	Title     string
	Summary   string
	Published string
	ReadMore  string
}

type PageData struct {
	Language        string
	Title           string
	Brand           string
	Home            string
	LanguageLabel   string
	HeroTitle       string
	HeroSubtitle    string
	PrimaryAction   string
	SecondaryAction string
	AboutTitle      string
	AboutText       string
	Projects        []Project
	BlogPosts       []BlogPost
	Email           string
	ProjectsText    string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	HomeHandlerWithConfig(config.Config{}, w, r)
}

func HomeHandlerWithConfig(cfg config.Config, w http.ResponseWriter, r *http.Request) {
	language := languageFromRequest(w, r)
	data := pageData(language)
	if cfg.GitHubUsername == "" {
		cfg.GitHubUsername = "falsisdev"
	}

	if cfg.GitHubUsername != "" {
		ghClient := github.New(github.Config{
			Username:   cfg.GitHubUsername,
			Token:      cfg.GitHubToken,
			HTTPClient: &http.Client{Timeout: 10 * time.Second},
		})
		if repos, err := ghClient.ListRepos(r.Context()); err == nil {
			data.Projects = githubReposToProjects(repos)
		}
	}

	if cfg.SanityProjectID != "" && cfg.SanityDataset != "" {
		sanityClient := sanity.New(sanity.Config{
			ProjectID:  cfg.SanityProjectID,
			Dataset:    cfg.SanityDataset,
			APIVersion: cfg.SanityAPIVersion,
			Token:      cfg.SanityToken,
		})
		if posts, err := sanityClient.ListPosts(r.Context()); err == nil {
			data.BlogPosts = sanityPostsToBlogPosts(posts)
		}
	}

	templateName := "base"
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
	projects := []Project{
		{Name: "Website Platform", Description: "Go, SSR and HTMX project for a lightweight personal product site.", Stack: "Go • HTMX • Tailwind", Link: "https://github.com/falsisdev/website"},
		{Name: "Realtime Activity Feed", Description: "Live GitHub and Sanity data delivered via SSE and server-side rendering.", Stack: "Go • SSE • API", Link: "https://github.com/falsisdev/website"},
	}
	blogPosts := []BlogPost{
		{Title: "Go ile minimalist web uygulamaları", Summary: "Server-side render ve HTMX ile hızlı, sade ve güçlü arayüzler inşa etmek.", Published: "2026 • Yazı", ReadMore: "Devamını oku"},
		{Title: "SSE ile gerçek zamanlı veri akışı", Summary: "Client tarafında yenileme yapmadan veri yayınlamak için uygun mimari örnekleri.", Published: "2026 • Yazı", ReadMore: "Devamını oku"},
	}

	if language == "tr" {
		return PageData{
			Language:        "tr",
			Title:           "Kişisel Web Sitesi",
			Brand:           "falsis.dev",
			Home:            "Ana Sayfa",
			LanguageLabel:   "Dil",
			HeroTitle:       "YTU'de EHM öğrencisi, kendince full-stack programlama ve web sistemleri üzerine odaklanan bir geliştirici.",
			HeroSubtitle:    "Kullanıcı deneyimi, veri akışı ve web altyapısı üzerine odaklanan minimal, performans odaklı çözümler geliştiriyorum.",
			PrimaryAction:   "Projeleri gör",
			SecondaryAction: "İletişime geç",
			AboutTitle:      "Hakkımda",
			AboutText:       "Yazılımı sadece teknoloji olarak değil, kullanıcıyla anlamlı bir ilişki kuran sistem olarak görüyorum. Arayüz, backend ve veri akışı arasında güçlü bir bütünlük kurup ürünleri daha net, daha hızlı ve daha güvenilir hale getiriyorum.",
			Email:           "falsis@proton.me",
			Projects:        projects,
			BlogPosts:       blogPosts,
			ProjectsText:    "Projeler",
		}
	}

	return PageData{
		Language:        "en",
		Title:           "Personal Website",
		Brand:           "falsis.dev",
		Home:            "Home",
		LanguageLabel:   "Language",
		HeroTitle:       "ECE undergraduate @ YTU, self-taught full-stack developer with a focus on web systems and backend architecture.",
		HeroSubtitle:    "I build minimal, high-performance experiences with a strong focus on backend architecture, user experience, and reliable data flows.",
		PrimaryAction:   "View projects",
		SecondaryAction: "Contact me",
		AboutTitle:      "About",
		AboutText:       "I see software as a system that creates meaningful relationships between users and technology. I aim to connect interfaces, backend services, and data flows into one coherent product experience that feels fast, clear, and dependable.",
		Email:           "falsis@proton.me",
		Projects:        projects,
		BlogPosts:       blogPosts,
		ProjectsText:    "Projects",
	}
}

func githubReposToProjects(repos []github.Repo) []Project {
	projects := make([]Project, 0, len(repos))
	for _, repo := range repos {
		if repo.Name == "" {
			continue
		} else if repo.Name == "falsisdev" {
			continue
		} else if repo.Name == "mangile" {
			projects = append(projects, Project{
				Name:        repo.Name,
				Description: repo.Description,
				Stack:       "Nuxt",
				Link:        repo.HTMLURL,
			})
		} else if repo.Name == "vessel" {
			projects = append(projects, Project{
				Name:        repo.Name,
				Description: repo.Description,
				Stack:       "Go • Swift",
				Link:        repo.HTMLURL,
			})
		} else if repo.Name == "website" {
			projects = append(projects, Project{
				Name:        repo.Name,
				Description: repo.Description,
				Stack:       "Go • HTMX",
				Link:        repo.HTMLURL,
			})
		} else {
			projects = append(projects, Project{
				Name:        repo.Name,
				Description: repo.Description,
				Stack:       repo.Language,
				Link:        repo.HTMLURL,
			})
		}
		continue
	}
	return projects
}

func sanityPostsToBlogPosts(posts []sanity.Post) []BlogPost {
	out := make([]BlogPost, 0, len(posts))
	for _, post := range posts {
		if post.Title == "" {
			continue
		}
		out = append(out, BlogPost{
			Title:     post.Title,
			Summary:   post.Summary,
			Published: normalizePublishedDate(post.Published),
			ReadMore:  "Read more",
		})
	}
	return out
}

func normalizePublishedDate(value string) string {
	if value == "" {
		return "Recently"
	}
	if strings.Contains(value, "T") {
		if t, err := time.Parse(time.RFC3339, value); err == nil {
			return t.Format("2006 • Jan")
		}
	}
	return value
}
