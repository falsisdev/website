package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	Username   string
	Token      string
	BaseURL    string
	HTTPClient *http.Client
}

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Language    string `json:"language"`
}

type Profile struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	AvatarURL   string `json:"avatar_url"`
	Bio         string `json:"bio"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	PublicRepos int    `json:"public_repos"`
}

type Client struct {
	config Config
}

func New(cfg Config) *Client {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.github.com"
	}
	return &Client{config: cfg}
}

func (c *Client) GetProfile(ctx context.Context) (*Profile, error) {
	if c.config.Username == "" {
		return nil, fmt.Errorf("github username is required")
	}

	url := fmt.Sprintf("%s/users/%s", strings.TrimRight(c.config.BaseURL, "/"), c.config.Username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if c.config.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.Token)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api returned status %d", res.StatusCode)
	}

	var profile Profile
	if err := json.NewDecoder(res.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

func (c *Client) ListRepos(ctx context.Context) ([]Repo, error) {
	if c.config.Username == "" {
		return nil, fmt.Errorf("github username is required")
	}

	url := fmt.Sprintf("%s/users/%s/repos", strings.TrimRight(c.config.BaseURL, "/"), c.config.Username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if c.config.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.Token)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api returned status %d", res.StatusCode)
	}

	var repos []Repo
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}
