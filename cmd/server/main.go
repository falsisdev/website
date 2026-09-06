package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/falsisdev/website/handlers"
	"github.com/falsisdev/website/internal/config"
	"github.com/falsisdev/website/internal/github"
	"github.com/falsisdev/website/internal/realtime"
	"github.com/falsisdev/website/internal/server"
)

func main() {
	if err := handlers.InitTemplates(); err != nil {
		slog.Error("template initialization failed", "error", err)
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		slog.Error("configuration load failed", "error", err)
		os.Exit(1)
	}

	bus := realtime.NewBus()
	poller := github.NewPoller(github.PollerConfig{
		Username: cfg.GitHubUsername,
		Token:    cfg.GitHubToken,
		Bus:      bus,
		Interval: 30 * time.Second,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go poller.Start(ctx)

	if err := poller.PollOnce(context.Background()); err != nil {
		slog.Warn("initial github poll failed", "error", err)
	}

	httpServer := &http.Server{
		Addr:              cfg.Host + ":" + cfg.Port,
		Handler:           server.NewMuxWithBus(cfg, bus),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		slog.Info("shutdown signal received, starting graceful shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			slog.Error("server shutdown failed", "error", err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	slog.Info("server starting", "address", httpServer.Addr, "environment", cfg.Environment)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped cleanly")
}
