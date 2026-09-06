package github

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/falsisdev/website/internal/realtime"
)

type PollerConfig struct {
	Username   string
	Token      string
	BaseURL    string
	HTTPClient *http.Client
	Interval   time.Duration
	Bus        *realtime.Bus
}

func NewPoller(cfg PollerConfig) *Poller {
	if cfg.Interval <= 0 {
		cfg.Interval = 30 * time.Second
	}
	return &Poller{config: cfg}
}

type Poller struct {
	config PollerConfig
}

func (p *Poller) Start(ctx context.Context) {
	if p.config.Bus == nil {
		return
	}
	if p.config.Username == "" {
		return
	}

	ticker := time.NewTicker(p.config.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := p.PollOnce(ctx); err != nil {
				slog.Warn("github poll failed", "error", err)
			}
		}
	}
}

func (p *Poller) PollOnce(ctx context.Context) error {
	if p.config.Username == "" {
		return nil
	}

	client := New(Config{
		Username:   p.config.Username,
		Token:      p.config.Token,
		BaseURL:    p.config.BaseURL,
		HTTPClient: p.config.HTTPClient,
	})

	profile, err := client.GetProfile(ctx)
	if err != nil {
		return err
	}

	payload := map[string]any{
		"login":        profile.Login,
		"name":         profile.Name,
		"avatar_url":   profile.AvatarURL,
		"bio":          profile.Bio,
		"followers":    profile.Followers,
		"following":    profile.Following,
		"public_repos": profile.PublicRepos,
	}

	if p.config.Bus != nil {
		p.config.Bus.Publish(realtime.Event{Type: "github_profile_updated", Data: payload})
	}
	return nil
}
