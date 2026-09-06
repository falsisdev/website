package sanity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	ProjectID  string
	Dataset    string
	APIVersion string
	Token      string
	BaseURL    string
	HTTPClient *http.Client
}

type Post struct {
	ID        string `json:"_id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Slug      string `json:"slug"`
	Published string `json:"publishedAt"`
}

type Client struct {
	config Config
}

type queryResponse struct {
	Result []struct {
		ID      string `json:"_id"`
		Title   string `json:"title"`
		Summary string `json:"summary"`
		Slug    struct {
			Current string `json:"current"`
		} `json:"slug"`
		Published string `json:"publishedAt"`
	} `json:"result"`
}

func New(cfg Config) *Client {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://<project-id>.api.sanity.io"
	}
	return &Client{config: cfg}
}

func (c *Client) ListPosts(ctx context.Context) ([]Post, error) {
	if c.config.ProjectID == "" || c.config.Dataset == "" {
		return nil, fmt.Errorf("sanity project id and dataset are required")
	}

	endpoint := fmt.Sprintf("%s/data/query/%s", strings.TrimRight(c.config.BaseURL, "/"), c.config.Dataset)
	if strings.Contains(c.config.BaseURL, "<project-id>") {
		endpoint = strings.Replace(endpoint, "https://<project-id>.api.sanity.io", fmt.Sprintf("https://%s.api.sanity.io", c.config.ProjectID), 1)
	}

	payload := map[string]string{
		"query": `*[_type == "post"] | order(publishedAt desc)[0..4]`,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.config.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.Token)
	}
	if c.config.APIVersion != "" {
		req.Header.Set("X-Sanity-API-Version", c.config.APIVersion)
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sanity api returned status %d", res.StatusCode)
	}

	var resp queryResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(resp.Result))
	for _, item := range resp.Result {
		posts = append(posts, Post{
			ID:        item.ID,
			Title:     item.Title,
			Summary:   item.Summary,
			Slug:      item.Slug.Current,
			Published: item.Published,
		})
	}
	return posts, nil
}
