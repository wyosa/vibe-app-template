package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"app/api/internal/app"
	"app/api/internal/config"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	cfg, err := config.Load()
	if err != nil {
		log.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	container, err := app.NewContainer(cfg, log)
	if err != nil {
		log.Error("failed to create container", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := container.Close(); err != nil {
			log.Error("failed to close container", "error", err)
		}
	}()

	application := app.New(container)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := application.Run(ctx); err != nil {
		log.Error("application error", "error", err)
		os.Exit(1)
	}
}
