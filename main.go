package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"streamgate/internal/config"
	"streamgate/internal/dependencies"
	"streamgate/internal/httpserver"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	deps := dependencies.New(cfg, log)

	srv := httpserver.New(deps)

	if err := srv.Start(ctx); err != nil {
		log.Error("server error", "error", err)
		os.Exit(1)
	}
}
