package main

import (
	"avito-tech-backend/internal/core"
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

	//TODO: init storage: pgx

	//TODO: init router: gin

	//TODO: run server

}
