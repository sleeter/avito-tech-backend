package main

import (
	"avito-tech-backend/internal/core"
	http_server "avito-tech-backend/internal/http-server"
	"avito-tech-backend/internal/pkg/config"
	"context"
	"log"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	loader := config.PrepareLoader(config.WithConfigPath("./config.yaml"))

	cfg, err := core.ParseConfig(loader)
	if err != nil {
		log.Fatalf("Failed to parse config: %s", err)
	}

	repository, err := core.NewRepository(ctx, cfg)
	if err != nil {
		log.Fatalf("Init repository: %s", err)
	}
	app := http_server.New(repository)

	if err := app.Start(ctx); err != nil {
		log.Fatalf(err.Error())
	}

}
